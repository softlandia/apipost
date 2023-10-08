package main

import (
	"apipost/model/user"
	"apipost/server"
	"context"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	usr := user.NewUser(10, "Имя", "Фамилия")
	fmt.Printf("usr id: :%s\n", usr.Id)

	fmt.Printf("start api post service\n")
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	cfg.Print()

	srv := &http.Server{
		Addr:    cfg.Listen,
		Handler: server.New(cfg.Listen),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("Server forced to shutdown, err: %s", err.Error())
		os.Exit(0)
	}

	log.Print("Shutting down server...")
}
