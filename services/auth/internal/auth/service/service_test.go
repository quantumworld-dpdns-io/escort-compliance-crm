package service

import (
	"testing"
	"time"
)

func TestGenerateTokenPair(t *testing.T) {
	svc := &Service{
		repo:      nil,
		cache:     nil,
		jwtSecret: "test-secret-for-unit-tests",
	}

	user := &User{
		ID:    "test-user-id-123",
		Email: "test@example.com",
		Role:  "client",
	}

	tokens, err := svc.generateTokenPair(user)
	if err != nil {
		t.Fatalf("generateTokenPair failed: %v", err)
	}

	if tokens.AccessToken == "" {
		t.Error("access token is empty")
	}
	if tokens.RefreshToken == "" {
		t.Error("refresh token is empty")
	}
	if tokens.TokenType != "Bearer" {
		t.Errorf("expected token type Bearer, got %s", tokens.TokenType)
	}
	if tokens.ExpiresIn != 900 {
		t.Errorf("expected expires_in 900, got %d", tokens.ExpiresIn)
	}
}

func TestValidateToken(t *testing.T) {
	svc := &Service{
		repo:      nil,
		cache:     nil,
		jwtSecret: "test-secret-for-unit-tests",
	}

	user := &User{
		ID:    "test-user-id-456",
		Email: "validate@example.com",
		Role:  "admin",
	}

	tokens, err := svc.generateTokenPair(user)
	if err != nil {
		t.Fatalf("generateTokenPair failed: %v", err)
	}

	claims, err := svc.validateToken(tokens.AccessToken)
	if err != nil {
		t.Fatalf("validateToken failed: %v", err)
	}

	if claims.UserID != "test-user-id-456" {
		t.Errorf("expected user_id test-user-id-456, got %s", claims.UserID)
	}
	if claims.Email != "validate@example.com" {
		t.Errorf("expected email validate@example.com, got %s", claims.Email)
	}
	if claims.Role != "admin" {
		t.Errorf("expected role admin, got %s", claims.Role)
	}
}

func TestValidateToken_InvalidSecret(t *testing.T) {
	svc1 := &Service{jwtSecret: "secret-1"}
	svc2 := &Service{jwtSecret: "secret-2"}

	user := &User{ID: "user-1", Email: "a@b.com", Role: "client"}
	tokens, _ := svc1.generateTokenPair(user)

	_, err := svc2.validateToken(tokens.AccessToken)
	if err == nil {
		t.Error("expected error validating token with wrong secret")
	}
}

func TestValidateToken_Expired(t *testing.T) {
	svc := &Service{jwtSecret: "test-secret"}

	claims := &Claims{
		UserID: "expired-user",
		Email:  "expired@test.com",
		Role:   "client",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
			Subject:   "expired-user",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString([]byte("test-secret"))

	_, err := svc.validateToken(tokenStr)
	if err == nil {
		t.Error("expected error for expired token")
	}
}

func TestHashToken(t *testing.T) {
	hash1 := hashToken("test-token")
	hash2 := hashToken("test-token")
	hash3 := hashToken("different-token")

	if hash1 != hash2 {
		t.Error("same input should produce same hash")
	}
	if hash1 == hash3 {
		t.Error("different inputs should produce different hashes")
	}
	if len(hash1) != 64 {
		t.Errorf("expected hash length 64, got %d", len(hash1))
	}
}
