package main

import (
	"context"
	"github.com/softlandia/hismap/pkg/config"
	"github.com/softlandia/hismap/pkg/mongo_db"
	"github.com/softlandia/hismap/repo/items"
	"github.com/softlandia/hismap/service"
	"github.com/softlandia/hismap/service/http_server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
)

func internalServer(logger *zerolog.Logger, port string) *http.Server {
	res := &http.Server{
		Addr:    port,
		Handler: http_server.InternalHandler(logger),
	}

	go func() {
		if err := res.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal().
				Err(err).
				Str("call", "ListenAndServe").
				Msg("internal server listen error")
		}
	}()
	return res
}

func mainServer(port string, svc *service.Service) *http.Server {
	res := &http.Server{
		Addr:    port,
		Handler: http_server.MainHandler(svc),
	}

	go func() {
		if err := res.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			svc.Logger.Fatal().
				Err(err).
				Str("call", "mainServer.ListenAndServe").
				Msg("server listen error")
		}
	}()
	return res
}

func adminServer(port string, svc *service.Service) *http.Server {
	res := &http.Server{
		Addr:    port,
		Handler: http_server.AdminHandler(svc),
	}

	go func() {
		if err := res.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			svc.Logger.Fatal().
				Err(err).
				Str("call", "adminServer.ListenAndServe").
				Msg("server listen error")
		}
	}()
	return res
}

func mainService(logger *zerolog.Logger, ctx context.Context, connect mongo_db.Connect, repos service.Repositories) *service.Service {
	svc := service.NewService(service.Config{
		Log:          logger,
		Ctx:          ctx,
		Db:           connect,
		Repositories: repos,
	})
	return svc
}

func logLevel(level string) zerolog.Level {
	res := zerolog.DebugLevel
	switch level {
	case "warn":
		res = zerolog.WarnLevel
	case "info":
		res = zerolog.InfoLevel
	}
	return res
}

// @title hismap server
// @version 0.0.1
// @description history map
// @contact.email softlandia@gmail.com
// @BasePath /api/v1
func main() {
	var cfg config.Configuration
	if err := envconfig.Process("", &cfg); err != nil {
		log.Printf("failed to load envconfig, error: %s", err.Error())
		return
	}
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger().Level(logLevel(cfg.LogLevel))

	ctx := context.Background()

	connect, err := mongo_db.New(ctx, cfg.MongoConnString, cfg.MongoDBName)
	if err != nil {
		logger.Fatal().Err(err).Msg("fail connect to mongodb")
	}
	if err := testConnect(connect); err != nil {
		logger.Fatal().Err(err).Msg("fail test connect to db")
	}
	logger.Info().Msg(">>> test db connect SUCCESS")

	repos := service.Repositories{
		Items: items.New(ctx, &logger, connect),
	}
	if err := testRepo(logger, repos); err != nil {
		return
	}
	logger.Info().Msg(">>> test repos SUCCESS")

	iServer := internalServer(&logger, cfg.ListenInternal)
	mService := mainService(&logger, ctx, connect, repos)
	mServer := mainServer(cfg.Listen, mService)
	aServer := adminServer(cfg.ListenAdmin, mService)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info().Msg("Shut down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := aServer.Shutdown(ctx); err != nil {
		logger.Fatal().
			Err(err).
			Msg("Admin server forced to shutdown")
	}
	if err := mServer.Shutdown(ctx); err != nil {
		logger.Fatal().
			Err(err).
			Msg("Main server forced to shutdown")
	}

	if err := iServer.Shutdown(ctx); err != nil {
		logger.Fatal().
			Err(err).
			Msg("Internal server forced to shutdown")
	}

	logger.Info().Msg("hismap server exiting")
}
