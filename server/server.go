package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

func Init() {
	router := newRouter()

	srv := http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 10)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	logrus.Info("Server shutting down in 0s ...")

	ctx, cancel := context.WithTimeout(context.Background(), 0*time.Second)
	defer cancel()

	<-ctx.Done()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server shut down with error: %s", err.Error())
	}

	logrus.Info("Server shut down")
}
