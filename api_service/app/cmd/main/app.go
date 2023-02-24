package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dexises/4room/api_service/internal/config"
	"github.com/dexises/4room/api_service/internal/user"
	"github.com/dexises/4room/api_service/pkg/logging"
)

func main() {
	logger := logging.GetLoggerInstance()
	logger.PrintInfo("logger initialized")

	cfg := config.GetConfigInstance()
	logger.PrintInfo("config initialized")

	router := http.NewServeMux()
	logger.PrintInfo("router initialized")

	handler := user.NewHandler()
	handler.Register(router)

	serve(router, logger, cfg)
}

func serve(router *http.ServeMux, logger *logging.Logger, cfg *config.Config) {
	server := &http.Server{
		Addr:         cfg.Port,
		Handler:      router,
		IdleTimeout:  time.Minute,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  25 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		logger.PrintInfo("caught signal:" + s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		logger.PrintInfo("completing background tasks" + server.Addr)

		shutdownError <- nil
	}()

	logger.PrintInfo("starting server" + server.Addr)

	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		logger.PrintFatal(err)
	}

	err = <-shutdownError
	if err != nil {
		logger.PrintError(err)
	}

	logger.PrintInfo("stopped server" + server.Addr)
}
