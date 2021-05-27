package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alseiitov/bookstore/internal/config"
	"github.com/alseiitov/bookstore/internal/handler"
	"github.com/alseiitov/bookstore/internal/repository"
	"github.com/alseiitov/bookstore/internal/server"
	"github.com/alseiitov/bookstore/internal/service"
	"github.com/alseiitov/bookstore/pkg/database"
)

func Run(configPath string) {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.GetConnection(
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
	)
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepositories(db)

	services := service.NewServices(service.ServiceDeps{
		Repos: repos,
	})

	handler := handler.NewHandler(services)

	server := server.NewServer(cfg, handler.Init())

	log.Println("app started")

	go func() {
		log.Fatal(server.Run())
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	log.Println("app shutting down")

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Printf("error occured on db connection close: %s", err.Error())
	}
}
