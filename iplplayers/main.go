package main

import (
	"context"
	"golangliveprojects/iplplayers/internal/db/mysql"
	"golangliveprojects/iplplayers/internal/handlers"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	DB = mysql.InitDB()

	// fmt.Println("DB:", DB)
	// set context
	// var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	// defer cancel()

	// // var playerData []response.PlayerResponse
	// // queries.ListQuery(ctx, &playerData, DB)

	// // fmt.Println("playerData:", playerData)
	// // fmt.Printf("playerData%#v:", playerData)
	router := handlers.SetupRouter(DB)
	serverPort := ":8080"
	listenAndServe(router, serverPort)
}

func listenAndServe(router *gin.Engine, port string) {
	log.Println("In listenAndServe start")
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		log.Printf("Listening on address: %s", port)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Printf("Shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Printf("Server exiting")
}
