package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	MQ       MQConfig
	Auth     AuthConfig
	Quantum  QuantumConfig
	ML       MLConfig
	Security SecurityConfig
}

type ServerConfig struct {
	Port         int
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type DatabaseConfig struct {
	URL            string
	MaxConns       int
	MinConns       int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

type RedisConfig struct {
	URL      string
	Password string
	DB       int
}

type MQConfig struct {
	RabbitMQ URL
	Kafka    URL
}

type URL string

type AuthConfig struct {
	JWTSecret          string
	JWTExpiry          time.Duration
	JWTRefreshExpiry   time.Duration
	OAuth2GitHubID     string
	OAuth2GitHubSecret string
	OAuth2GoogleID     string
	OAuth2GoogleSecret string
}

type QuantumConfig struct {
	Enabled    bool
	Algorithm  string
	OptEnabled bool
	MLDisabled bool
	QRNGEnabled bool
}

type MLConfig struct {
	OpenAIKey    string
	AnthropicKey string
	OllamaURL    string
}

type SecurityConfig struct {
	TurnstileSiteKey   string
	TurnstileSecretKey string
	WAFEnabled         bool
	PQCEnabled         bool
}

func Load() (*Config, error) {
	cfg := &Config{}

	cfg.Server = ServerConfig{
		Port:         getEnvInt("PORT", 8080),
		Host:         getEnv("HOST", "0.0.0.0"),
		ReadTimeout:  getEnvDuration("READ_TIMEOUT", 30*time.Second),
		WriteTimeout: getEnvDuration("WRITE_TIMEOUT", 30*time.Second),
		IdleTimeout:  getEnvDuration("IDLE_TIMEOUT", 120*time.Second),
	}

	cfg.Database = DatabaseConfig{
		URL:            getEnv("DATABASE_URL", "postgres://localhost:5432/escort_crm"),
		MaxConns:       getEnvInt("DATABASE_MAX_CONNS", 25),
		MinConns:       getEnvInt("DATABASE_MIN_CONNS", 5),
		ConnMaxLifetime: getEnvDuration("DATABASE_CONN_MAX_LIFETIME", 5*time.Minute),
		ConnMaxIdleTime: getEnvDuration("DATABASE_CONN_MAX_IDLE_TIME", 5*time.Minute),
	}

	cfg.Redis = RedisConfig{
		URL:      getEnv("REDIS_URL", "redis://localhost:6379"),
		Password: getEnv("REDIS_PASSWORD", ""),
		DB:       getEnvInt("REDIS_DB", 0),
	}

	cfg.Auth = AuthConfig{
		JWTSecret:          getEnv("JWT_SECRET", "change-me"),
		JWTExpiry:          getEnvDuration("JWT_EXPIRY", 15*time.Minute),
		JWTRefreshExpiry:   getEnvDuration("JWT_REFRESH_EXPIRY", 168*time.Hour),
		OAuth2GitHubID:     getEnv("OAUTH2_GITHUB_CLIENT_ID", ""),
		OAuth2GitHubSecret: getEnv("OAUTH2_GITHUB_CLIENT_SECRET", ""),
		OAuth2GoogleID:     getEnv("OAUTH2_GOOGLE_CLIENT_ID", ""),
		OAuth2GoogleSecret: getEnv("OAUTH2_GOOGLE_CLIENT_SECRET", ""),
	}

	cfg.Quantum = QuantumConfig{
		Enabled:    getEnvBool("PQC_ENABLED", true),
		Algorithm:  getEnv("PQC_ALGORITHM", "kyber768"),
		OptEnabled: getEnvBool("QUANTUM_OPT_ENABLED", true),
		MLDisabled: getEnvBool("QUANTUM_ML_ENABLED", true),
		QRNGEnabled: getEnvBool("QRNG_ENABLED", false),
	}

	cfg.ML = MLConfig{
		OpenAIKey:    getEnv("OPENAI_API_KEY", ""),
		AnthropicKey: getEnv("ANTHROPIC_API_KEY", ""),
		OllamaURL:    getEnv("OLLAMA_BASE_URL", "http://localhost:11434"),
	}

	cfg.Security = SecurityConfig{
		TurnstileSiteKey:   getEnv("TURNSTILE_SITE_KEY", ""),
		TurnstileSecretKey: getEnv("TURNSTILE_SECRET_KEY", ""),
		WAFEnabled:         getEnvBool("WAF_ENABLED", true),
		PQCEnabled:         getEnvBool("PQC_ENABLED", true),
	}

	return cfg, nil
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return defaultVal
}

func getEnvBool(key string, defaultVal bool) bool {
	if val := os.Getenv(key); val != "" {
		if b, err := strconv.ParseBool(val); err == nil {
			return b
		}
	}
	return defaultVal
}

func getEnvDuration(key string, defaultVal time.Duration) time.Duration {
	if val := os.Getenv(key); val != "" {
		if d, err := time.ParseDuration(val); err == nil {
			return d
		}
	}
	return defaultVal
}

func (c *Config) Validate() error {
	var errs []string

	if c.Database.URL == "" {
		errs = append(errs, "DATABASE_URL is required")
	}
	if c.Auth.JWTSecret == "" || c.Auth.JWTSecret == "change-me" {
		errs = append(errs, "JWT_SECRET must be set")
	}
	if c.Server.Port < 1 || c.Server.Port > 65535 {
		errs = append(errs, fmt.Sprintf("invalid PORT: %d", c.Server.Port))
	}

	if len(errs) > 0 {
		return fmt.Errorf("config validation failed: %s", strings.Join(errs, "; "))
	}
	return nil
}
