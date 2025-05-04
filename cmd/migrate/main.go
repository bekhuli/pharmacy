package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bekhuli/pharmacy/internal/common"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const migrationsDir = "file://migrations"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("missing command: up | down")
	}
	command := os.Args[1]

	cfg := common.DBEnv

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.PublicHost,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
	)

	m, err := migrate.New(migrationsDir, dbURL)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}

	switch command {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("migration up failed: %v", err)
		}
		log.Println("all migrations applied successfully")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("migration down failed: %v", err)
		}
		log.Println("all migrations reverted successfully")
	default:
		log.Fatalf("unknown command: %s (expected: up | down)", command)
	}
}
