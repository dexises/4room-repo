package graceful

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dexises/4room/api_service/pkg/logging"
)

func Graceful(logger *logging.Logger, server *http.Server, shutdownError chan error) {
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
}
