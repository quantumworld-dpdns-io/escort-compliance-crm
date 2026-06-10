package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/auth/internal/auth/repository"
	"github.com/quantumworld-dpdns-io/escort-compliance-crm/services/shared/pkg/cache"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      *repository.Repository
	cache     *cache.Cache
	jwtSecret string
}

func New(repo *repository.Repository, cache *cache.Cache, jwtSecret string) *Service {
	return &Service{repo: repo, cache: cache, jwtSecret: jwtSecret}
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func (s *Service) Register(ctx context.Context, email, password, name, role string) (*repository.User, error) {
	exists, err := s.repo.ExistsByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("check email: %w", err)
	}
	if exists {
		return nil, errors.New("email already registered")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	if role == "" {
		role = "client"
	}

	user, err := s.repo.Create(ctx, email, string(hash), name, role)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	_ = s.repo.CreateAuditLog(ctx, user.ID, "register", "user", user.ID, "", "", nil)
	return user, nil
}

func (s *Service) Login(ctx context.Context, email, password, ip, userAgent string) (*TokenPair, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if user.LockedUntil != nil && user.LockedUntil.After(time.Now()) {
		return nil, errors.New("account locked, try again later")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		_ = s.repo.IncrementFailedLogin(ctx, email)
		return nil, errors.New("invalid credentials")
	}

	tokens, err := s.generateTokenPair(user)
	if err != nil {
		return nil, fmt.Errorf("generate tokens: %w", err)
	}

	tokenHash := hashToken(tokens.RefreshToken)
	expiresAt := time.Now().Add(7 * 24 * time.Hour)
	_ = s.repo.CreateSession(ctx, user.ID, tokenHash, ip, userAgent, expiresAt)
	_ = s.repo.UpdateLastLogin(ctx, user.ID)
	_ = s.repo.CreateAuditLog(ctx, user.ID, "login", "user", user.ID, ip, userAgent, nil)

	return tokens, nil
}

func (s *Service) RefreshToken(ctx context.Context, refreshToken string) (*TokenPair, error) {
	claims, err := s.validateToken(refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	tokenHash := hashToken(refreshToken)
	_ = s.repo.RevokeSession(ctx, tokenHash)

	user, err := s.repo.FindByID(ctx, claims.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	tokens, err := s.generateTokenPair(user)
	if err != nil {
		return nil, fmt.Errorf("generate tokens: %w", err)
	}

	return tokens, nil
}

func (s *Service) Logout(ctx context.Context, refreshToken string) error {
	tokenHash := hashToken(refreshToken)
	return s.repo.RevokeSession(ctx, tokenHash)
}

func (s *Service) ValidateAccessToken(tokenString string) (*Claims, error) {
	return s.validateToken(tokenString)
}

func (s *Service) GetUser(ctx context.Context, userID string) (*repository.User, error) {
	return s.repo.FindByID(ctx, userID)
}

func (s *Service) generateTokenPair(user *repository.User) (*TokenPair, error) {
	accessClaims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   user.ID,
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessStr, err := accessToken.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, err
	}

	refreshClaims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   user.ID,
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshStr, err := refreshToken.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessStr,
		RefreshToken: refreshStr,
		ExpiresIn:    900,
		TokenType:    "Bearer",
	}, nil
}

func (s *Service) validateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func hashToken(token string) string {
	h := sha256.Sum256([]byte(token))
	return hex.EncodeToString(h[:])
}
