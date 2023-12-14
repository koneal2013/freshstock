package cmd

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/koneal2013/freshstock/internal/api"
	"github.com/koneal2013/freshstock/internal/model"
	"github.com/koneal2013/freshstock/internal/store"
)

func main() {
	myStore := store.NewProduceStore()
	handlers := api.NewHandlers(myStore)
	engine := api.SetupRoutes(handlers)

	err := model.RegisterValidator()
	if err != nil {
		log.Fatal(err.Error())
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	// Initializing the server in a goroutine so that it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
