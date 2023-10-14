package main

import (
	"apipost/repositories/orders_repo"
	"apipost/repositories/users_repo"
	"apipost/server"
	"apipost/service"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	fmt.Printf("start api post service\n")
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	cfg.Print()

	conn, err := pgx.Connect(ctx, cfg.PgConnectUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	repos := service.Repositories{
		Users:  users_repo.New(ctx, conn),
		Orders: orders_repo.New(ctx, conn),
	}
	srvc := service.New(repos)

	srv := &http.Server{
		Addr:    cfg.Listen,
		Handler: server.New(cfg.Listen, srvc),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown, err: %s", err.Error())
		os.Exit(0)
	}

	log.Print("Shutting down server...")
}
