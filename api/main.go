package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	// "os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/putitaT/skill-api-kafka/api/database"
	"github.com/putitaT/skill-api-kafka/api/skill"
)

func init() {
	db := database.ConnectDB()
	database.CreateTable(db)
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	r := gin.Default()

	db := database.ConnectDB()
	skillRepository := skill.NewRepository(db)
	skillHandler := skill.NewHandler(skillRepository)
	skill.Router(r, skillHandler)

	srv := http.Server{
		Addr:    ":" + "8090",
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
