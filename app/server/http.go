package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"htemplx/app/conf"
	"htemplx/pkg/logger"
)

func Run() {
	// get environment config
	cfg := conf.NewConfig()

	// set default slog logger
	logger.SetDefault(logger.GetLogger(cfg.LogLevel))

	// define server properties
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.HttpHost, cfg.HttpPort),
		Handler:      setupRouter(),
		ReadTimeout:  cfg.HttpReadTimeout,
		WriteTimeout: cfg.HttpWriteTimeout,
	}

	// context for managing server lifecycle
	serverCtx, serverCtxCancel := context.WithCancel(context.Background())

	// channel to listen for interrupt/quit signals
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		slog.Info("shutdown signal received, initiating graceful shutdown...")

		// context for graceful shutdown with a timeout
		shutdownCtx, shutdownCtxCancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer shutdownCtxCancel()

		// enforce shutdown timeout
		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				slog.Error("graceful shutdown timedout, forcing exit")
				os.Exit(1)
			}
		}()

		// trigger graceful shutdown
		if err := server.Shutdown(shutdownCtx); err != nil {
			slog.Error("error during shutdown", "error", err)
			os.Exit(1)
		}
		serverCtxCancel()
	}()

	// start the server
	slog.Info("starting http server...", "addr", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("failed to start server", "error", err)
		os.Exit(1)
	}

	// wait for the server context to be cancelled (i.e., shutdown complete)
	<-serverCtx.Done()
	slog.Info("server gracefully stopped...")
}
