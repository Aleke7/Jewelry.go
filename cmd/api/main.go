package main

import (
	"context"
	"database/sql"
	"flag"
	_ "github.com/lib/pq"
	"jewelry.abgdrv.com/internal/data"
	"jewelry.abgdrv.com/internal/jsonlog"
	"jewelry.abgdrv.com/internal/mailer"
	"os"
	"strings"
	"sync"
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
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
	cors struct {
		trustedOrigins []string
	}
}

// Application struct to hold HTTP handlers, helpers, middleware
type application struct {
	config config
	logger *jsonlog.Logger
	models data.Models
	mailer mailer.Mailer
	wg     sync.WaitGroup
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

	flag.Float64Var(&cfg.limiter.rps,
		"limiter-rps",
		2,
		"Rate limiter maximum requests per second")
	flag.IntVar(&cfg.limiter.burst,
		"limiter-burst",
		4,
		"Rate limiter maximum burst")
	flag.BoolVar(&cfg.limiter.enabled,
		"limiter-enabled",
		true,
		"Enable rate limiter")

	flag.StringVar(&cfg.smtp.host,
		"smtp-host",
		"sandbox.smtp.mailtrap.io",
		"SMTP host")
	flag.IntVar(&cfg.smtp.port,
		"smtp-port",
		25,
		"SMTP port")
	flag.StringVar(&cfg.smtp.username,
		"smtp-username",
		"f144a7aa352bc4",
		"SMTP username")
	flag.StringVar(&cfg.smtp.password,
		"smtp-password",
		"ddf324cb70a817",
		"SMTP password")
	flag.StringVar(&cfg.smtp.sender,
		"smtp-sender",
		"watch.me <no-reply@watch.me.abgdrv.com>",
		"SMTP sender")

	flag.Func("cors-trusted-origins",
		"Trusted CORS origins (space separated)",
		func(val string) error {
			cfg.cors.trustedOrigins = strings.Fields(val)
			return nil
		})

	flag.Parse()

	// Initialization of logger (recording information about the execution of an application)
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	defer db.Close()

	logger.PrintInfo("database connection pool established", nil)

	// Declare instance of application
	app := application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
		mailer: mailer.New(cfg.smtp.host,
			cfg.smtp.port,
			cfg.smtp.username,
			cfg.smtp.password,
			cfg.smtp.sender),
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}

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
