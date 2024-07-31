package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/cristiano-pacheco/pingo/internal/application/usecase/user/activateuseruc"
	"github.com/cristiano-pacheco/pingo/internal/application/usecase/user/authenticateuseruc"
	"github.com/cristiano-pacheco/pingo/internal/application/usecase/user/createuseruc"
	"github.com/cristiano-pacheco/pingo/internal/application/usecase/user/resetpassworduc"
	"github.com/cristiano-pacheco/pingo/internal/application/usecase/user/sendresetpasswordemailuc"
	"github.com/cristiano-pacheco/pingo/internal/application/usecase/user/updatepassworduc"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/configdm"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/keydm"
	"github.com/cristiano-pacheco/pingo/internal/domain/service/hashds"
	"github.com/cristiano-pacheco/pingo/internal/infra/database/repository/userrepo"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/handler/pinghandler"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/handler/user/activateuserhandler"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/handler/user/authenticateuserhandler"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/handler/user/createuserhandler"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/handler/user/resetpasswordhandler"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/handler/user/sendresetpasswordemailhandler"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/handler/user/updatepasswordhandler"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/middleware/authmw"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/middleware/loggermw"
	"github.com/cristiano-pacheco/pingo/internal/infra/http/response"
	"github.com/cristiano-pacheco/pingo/internal/infra/mailer/mailertemplate"
	"github.com/cristiano-pacheco/pingo/internal/infra/mailer/smtpmailer"
	"github.com/cristiano-pacheco/pingo/internal/infra/service/tokensvc"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"

	"github.com/go-mail/mail/v2"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

type config struct {
	port            int
	env             string
	apiBaseURL      string
	frontEndBaseURL string
	db              struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
	limiter struct {
		enabled bool
		rps     float64
		burst   int
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
	privateKeyPath string
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|staging|production)")

	defaultDSN := "postgres://postgres:123456789@127.0.0.1/pingo?sslmode=disable"
	flag.StringVar(&cfg.db.dsn, "db-dsn", defaultDSN, "PostgreSQL DSN")

	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.DurationVar(&cfg.db.maxIdleTime, "db-max-idle-time", 15*time.Minute, "PostgreSQL max connection idle time")

	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")
	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")

	defaultBaseURL := fmt.Sprintf("http://localhost:%d", cfg.port)
	flag.StringVar(&cfg.apiBaseURL, "api-base-url", defaultBaseURL, "API Base URL")
	flag.StringVar(&cfg.frontEndBaseURL, "frontend-base-url", "http://localhost:3000", "FrontEnd Base URL")

	flag.StringVar(&cfg.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "SMTP host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 2525, "SMTP port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", "dd7ff882c024ab", "SMTP username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", "3e2207712f6eef", "SMTP password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", "Pingo <no-reply@pingo.com>", "SMTP sender")

	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		cfg.cors.trustedOrigins = strings.Fields(val)
		return nil
	})

	flag.StringVar(&cfg.privateKeyPath, "private-key-path", "private.pem", "Private key path")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// -------------------------------------------------------------------------
	// Load the .env file

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// -------------------------------------------------------------------------
	// Private key load

	pemData, err := os.ReadFile(cfg.privateKeyPath)
	if err != nil {
		log.Fatalf("error loading private key file: %s", err)
	}

	key, err := keydm.New(pemData)
	if err != nil {
		log.Fatal(err)
	}

	// -------------------------------------------------------------------------
	// Create the configuration domain value object

	configVo, err := configdm.New(cfg.env, cfg.apiBaseURL, cfg.frontEndBaseURL)
	if err != nil {
		log.Fatalf("Error creating the config value object: %s", err)
	}

	// -------------------------------------------------------------------------
	// Connect to the database

	db, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("database connection pool established")

	// -------------------------------------------------------------------------
	// Run database migrations

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	m.Up()

	// -------------------------------------------------------------------------
	// Repository Creation

	userRepository := userrepo.New(db)

	// -------------------------------------------------------------------------
	// Service Creation

	hashService := hashds.New()

	defaultIssueName := "pingo"
	jwtParser := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Name}))
	tokenService := tokensvc.New(userRepository, key, jwtParser, defaultIssueName)

	// -------------------------------------------------------------------------
	// Gateways Creation
	dialer := mail.NewDialer(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password)
	smtpMailerGW := smtpmailer.New(dialer, cfg.smtp.sender)

	mailerTemplate := mailertemplate.MailerTemplate{}

	// -------------------------------------------------------------------------
	// UseCases Creation

	createUserMapper := createuseruc.NewMapper(hashService)
	createUserUseCase := createuseruc.New(userRepository, smtpMailerGW, mailerTemplate, configVo, createUserMapper)
	sendResetPasswordEmailUseCase := sendresetpasswordemailuc.New(userRepository, smtpMailerGW, mailerTemplate, hashService, configVo)
	resetPasswordUseCase := resetpassworduc.New(userRepository, hashService)
	activateUserUseCase := activateuseruc.New(userRepository)
	authenticateUserUseCase := authenticateuseruc.New(tokenService, userRepository, *hashService)
	updatePasswordUseCase := updatepassworduc.New(userRepository, *hashService)

	// -------------------------------------------------------------------------
	// Handlers Creation

	pingHandler := pinghandler.New()
	createUserHandler := createuserhandler.New(createUserUseCase)
	activateUserHandler := activateuserhandler.New(activateUserUseCase)
	sendResetPasswordEmailHandler := sendresetpasswordemailhandler.New(sendResetPasswordEmailUseCase)
	resetPasswordHandler := resetpasswordhandler.New(resetPasswordUseCase)
	authenticateUserHandler := authenticateuserhandler.New(authenticateUserUseCase)
	updatePasswordHandler := updatepasswordhandler.New(updatePasswordUseCase)

	// -------------------------------------------------------------------------
	// Middlewares

	authMiddleware := authmw.New(tokenService)

	// -------------------------------------------------------------------------
	// Routes registration

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(loggermw.AddLoggerToContextMiddleware(logger))

	router.NotFound(response.NotFoundResponse)
	router.MethodNotAllowed(response.MethodNotAllowedResponse)

	// public endpoints
	router.Group(func(r chi.Router) {
		// user
		router.Post("/api/v1/users", createUserHandler.Execute)
		router.Post("/api/v1/users/activate", activateUserHandler.Execute)
		router.Post("/api/v1/users/reset-password", sendResetPasswordEmailHandler.Execute)
		router.Put("/api/v1/users/reset-password", resetPasswordHandler.Execute)
		router.Post("/api/v1/users/auth", authenticateUserHandler.Execute)
	})

	// protected endpoints
	router.Group(func(r chi.Router) {
		r.Use(authMiddleware.Authenticate)
		r.Get("/api/v1/ping", pingHandler.Execute)

		r.Put("/api/v1/users/password", updatePasswordHandler.Execute)
	})

	// -------------------------------------------------------------------------
	// Start the webserver

	err = startWebServer(router, &cfg, logger)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	db.SetConnMaxIdleTime(cfg.db.maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func startWebServer(mux *chi.Mux, cfg *config, logger *slog.Logger) error {
	var wg sync.WaitGroup

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		logger.Info("caught signal", "signal", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		logger.Info("completing background tasks", "addr", srv.Addr)

		wg.Wait()
		shutdownError <- nil
	}()

	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	logger.Info("stopped server", "addr", srv.Addr)

	return nil
}
