package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/adsyandex/otus_shool/todo/internal/api"
	"github.com/adsyandex/otus_shool/todo/internal/config"
	"github.com/adsyandex/otus_shool/todo/internal/logger"
	"github.com/adsyandex/otus_shool/todo/internal/service"
	postgresstorage "github.com/adsyandex/otus_shool/todo/internal/storage/postgres"
	"github.com/golang-migrate/migrate/v4"
	postgresdriver "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log := logger.New("error")
		log.Fatal("Failed to load config: ", err)
	}

	log := logger.New(cfg.Log.Level)
	log.Info("Starting application version 1.0.0")

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User,
		cfg.Postgres.Password, cfg.Postgres.DBName, cfg.Postgres.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL: ", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping PostgreSQL: ", err)
	}

	if err := runMigrations(db); err != nil {
		log.Fatal("Database migrations failed: ", err)
	}

	storage, err := postgresstorage.NewPostgresStorage(connStr)
	if err != nil {
		log.Fatal("Failed to initialize storage: ", err)
	}

	taskService := service.NewTaskService(storage)
	router := api.NewRouter(taskService, log)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(cfg.Server.Port),
		Handler: router,
	}

	go func() {
		log.Info("Server starting on port ", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("HTTP server crashed: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server shutdown failed: ", err)
	}

	log.Info("Server stopped gracefully")
}

func runMigrations(db *sql.DB) error {
	driver, err := postgresdriver.WithInstance(db, &postgresdriver.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	fsrc, err := (&file.File{}).Open("file://migrations")
	if err != nil {
		return fmt.Errorf("failed to open migrations: %w", err)
	}

	m, err := migrate.NewWithInstance("file", fsrc, "postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
