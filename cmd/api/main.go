package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"jewelry.abgdrv.com/internal/data"
	"log"
	"net/http"
	"os"
	"time"
)

// Application version_number
const version = "1.0.0"

// Configuration settings
type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

// Application struct to hold HTTP handlers, helpers, middleware
type application struct {
	config config
	logger *log.Logger
	models data.Models
}

func main() {

	var cfg config

	// Parse configuration
	flag.IntVar(&cfg.port, "port",
		4000,
		"API server port")
	flag.StringVar(&cfg.env, "env",
		"development",
		"Environment (development|staging|production)")

	flag.StringVar(&cfg.db.dsn, "db-dsn",
		"postgres://watch_admin:watch@localhost/watch_database?sslmode=disable",
		"PostgreSQL DSN")

	flag.IntVar(&cfg.db.maxOpenConns,
		"db-max-open-conns",
		25,
		"PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns,
		"db-max-idle-conns",
		25,
		"PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime,
		"db-max-idle-time",
		"15m",
		"PostgreSQL max connection idle time")

	flag.Parse()

	// Initialization of logger (recording information about the execution of an application)
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	logger.Printf("database connection pool established")

	// Declare instance of application
	app := application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	// Declare HTTP server
	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start HTTP server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)

}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
