package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/putitaT/skill-api-kafka/api/skill"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	r := gin.Default()

	// database.CreateTable()
	skill.Producer()

	srv := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}

	closedChan := make(chan struct{})

	go gracefully(ctx, &srv, closedChan)

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}

	<-closedChan
	fmt.Println("bye")
}

func gracefully(ctx context.Context, srv *http.Server, clochan chan struct{}) {
	<-ctx.Done()
	fmt.Println("shutting down....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Println(err)
		}
	}

	close(clochan)
}
