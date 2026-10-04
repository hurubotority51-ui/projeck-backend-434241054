package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() error {
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		return fmt.Errorf("DATABASE_URL belum diatur")
	}

	var err error

	DB, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		return fmt.Errorf("gagal membuat koneksi database: %w", err)
	}

	if err := DB.Ping(context.Background()); err != nil {
		return fmt.Errorf("gagal terhubung ke database: %w", err)
	}

	fmt.Println("Database PostgreSQL berhasil terhubung")

	return nil
}
