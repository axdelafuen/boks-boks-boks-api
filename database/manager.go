package database

import (
  "fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Manager struct {
  db *gorm.DB
}

func NewManager() (*Manager, error) {
  db, err := initializeConnection()
  if err != nil {
    return nil, err
  }

  return &Manager{db: db}, nil
}

func (m *Manager) GetDB() *gorm.DB {
  return m.db
}

func (m *Manager) Close() error {
  sqlDB, err := m.db.DB()
  if err != nil {
    return err
  }

  return sqlDB.Close()
}

func initializeConnection() (*gorm.DB, error) {
  host := getEnvOrDefault("DB_HOST", "localhost")
	port := getEnvOrDefault("DB_PORT", "5432")
	user := getEnvOrDefault("DB_USER", "boks-dev")
	password := getEnvOrDefault("DB_PASSWORD", "123456")
	dbname := getEnvOrDefault("DB_NAME", "boks_db")

  dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  config := &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
  }

  db, err := gorm.Open(postgres.Open(dsn), config)
  if err != nil {
    return nil, fmt.Errorf("failed to connect to database: %w", err)
  }

  sqlDB, err := db.DB()
  if err != nil {
    return nil, err
  }

  sqlDB.SetMaxIdleConns(10)
  sqlDB.SetMaxOpenConns(100)

  return db, nil
}

func getEnvOrDefault(key, defaultValue string) string {
  if value := os.Getenv(key); value != "" {
    return value
  }

  return defaultValue
}
