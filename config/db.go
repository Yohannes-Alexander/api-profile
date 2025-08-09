package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// DBConfig holds the configuration values
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// LoadDBConfig loads database config from environment or fallback values
func LoadDBConfig() DBConfig {
	return DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "yourpass"),
		DBName:   getEnv("DB_NAME", "yourdb"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

// ConnectDB returns a database connection
func ConnectDB() (*sql.DB, error) {
	cfg := LoadDBConfig()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println("Failed to open database:", err)
		return nil, err
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		log.Println("Failed to connect to database:", err)
		return nil, err
	}

	log.Println("Successfully connected to the database")
	return db, nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
