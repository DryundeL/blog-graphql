package db

import (
	"errors"
	"github.com/DryundeL/blog-graphql/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"log"
)

func RunMigrations(cfg *config.Config) {
	m, err := migrate.New("file://migrations", cfg.DBConn)
	if err != nil {
		log.Fatalf("Ошибка миграций: %v", err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Ошибка применения миграций: %v", err)
	}
	log.Println("Миграции применены")
}
