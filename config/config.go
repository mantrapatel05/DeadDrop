package config

import(
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"github.com/joho/godotenv"
)

type Config struct {
	EncryptionKey []byte
	Port string
	DatabaseURL string
}

func LoadConfig() (*Config, error){
	godotenv.Load()

	keyHex := strings.TrimSpace(os.Getenv("ENCRYPTION_KEY"))
	if keyHex == "" {
		return nil, fmt.Errorf("ENCRYPTION_KEY is not set")
	}
	if len(keyHex) != 64 {
		return nil, fmt.Errorf("ENCRYPTION_KEY must be 64 characters long")
	}
	key,err := hex.DecodeString(keyHex)
	if err != nil{
		return nil, fmt.Errorf("invalid ENCRYPTION_KEY: %w", err)
	}
	if len(key) != 32 {
		return nil, fmt.Errorf("ENCRYPTION_KEY must be 32 bytes long")
	}

	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		port = "8080"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	databaseURL := strings.TrimSpace(os.Getenv("DATABASE_URL"))
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	return &Config{
		EncryptionKey: key,
		Port: port,
		DatabaseURL: databaseURL,
	}, nil
}