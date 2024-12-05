package app

import (
	"context"
	"github.com/jmoiron/sqlx"
	"houseService/config"
	"houseService/internal/adapter/dbs/postrges"
	httpApi "houseService/internal/handler/http"
	"houseService/internal/service"
	"houseService/pkg/auth"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	log             *slog.Logger
	db              *sqlx.DB
	httpServer      *http.Server
	serviceProvider *ServiceProvider
	tokenManager    *auth.Manager
}

func NewApp() (*App, error) {
	a := &App{}

	err := a.initDeps()
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() {
	go a.runHttpServer()

	a.log.Info("server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		a.log.Error("failed to stop server", slog.String("error", err.Error()))
	}

	a.log.Info("server stopped")

	if err := a.db.Close(); err != nil {
		a.log.Error("failed to stop storage", slog.String("error", err.Error()))
	}

}

func (a *App) initDeps() error {
	inits := []func() error{
		a.initConfig,
		a.initLogger,
		a.initDb,
		a.initTokenManager,
		a.initServiceProvider,
		a.initHttpServer,
	}

	for _, f := range inits {
		err := f()
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig() error {
	err := config.Load(".env")
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initLogger() error {
	a.log = slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	return nil
}

func (a *App) initDb() error {
	cfg, err := postrges.NewConfig()
	if err != nil {
		return err
	}

	db, err := postrges.New(cfg)
	if err != nil {
		a.log.Error("failed to init storage", slog.String("error", err.Error()))
		return err
	}

	a.db = db

	return nil
}

func (a *App) initTokenManager() error {
	cfg, err := service.NewConfig()
	if err != nil {
		return err
	}
	a.tokenManager, err = auth.NewManager(cfg.SigningKey, cfg.AccessTokenTTL)
	if err != nil {
		a.log.Error("failed to init token manager", slog.String("error", err.Error()))
		return err
	}

	return nil
}

func (a *App) initServiceProvider() error {
	a.serviceProvider = NewServiceProvider(a.log, a.db, a.tokenManager)
	return nil
}

func (a *App) initHttpServer() error {
	a.serviceProvider.RegisterControllers()

	cfg, err := httpApi.NewConfig()
	if err != nil {
		return err
	}

	srv := httpApi.NewServer(cfg, a.serviceProvider.HttpRouter())

	a.httpServer = srv
	return nil
}

func (a *App) runHttpServer() {
	if err := a.httpServer.ListenAndServe(); err != nil {
		a.log.Error("failed to start server", slog.String("error", err.Error()))
	}
}
