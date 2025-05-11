package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bekhuli/pharmacy/internal/common"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const migrationsDir = "file://migrations"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("missing command: up | down | force [version]")
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
	case "force":
		if len(os.Args) < 3 {
			log.Fatal("missing version for force command")
		}
		version, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("invalid version number: %v", err)
		}
		if err := m.Force(version); err != nil {
			log.Fatalf("force command failed: %v", err)
		}
		log.Printf("successfully forced version to %d", version)
	default:
		log.Fatalf("unknown command: %s (expected: up | down | force)", command)
	}
}
