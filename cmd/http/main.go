package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/brunabbelini/quicknotes/internal/mailer"
	"github.com/gorilla/csrf"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config := loadConfig()

	slog.SetDefault(newLogger(os.Stderr, config.GetLevelLog()))

	dbpool, err := pgxpool.New(context.Background(), config.DBConnURL)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info("Conexão com o banco estabelecida com sucesso")

	defer dbpool.Close()

	slog.Info(fmt.Sprintf("Servidor rodando na porta %s\n", config.ServerPort))

	// testando o envio de email
	mailPort, _ := strconv.Atoi(config.MailPort)
	mailService := mailer.NewSMTPMailService(mailer.SMTPConfig{
		Host:     config.MailHost,
		Port:     mailPort,
		Username: config.MailUsername,
		Password: config.MailPassword,
		From:     config.MailFrom,
	})

	sessionManager := scs.New()
	sessionManager.Lifetime = time.Hour
	sessionManager.Store = pgxstore.New(dbpool)
	pgxstore.NewWithCleanupInterval(dbpool, 30*time.Second)

	csrfMiddleware := csrf.Protect(
		[]byte(config.CSRFKey),
		csrf.TrustedOrigins([]string{
			"localhost:5000",
			"127.0.0.1:5000",
			"localhost:8443",
		}),
	)

	mux := LoadRoutes(sessionManager, mailService, dbpool)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), sessionManager.LoadAndSave(csrfMiddleware(mux))); err != nil {
		panic(err)
	}
}
