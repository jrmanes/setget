package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var (
	// port to the http server
	httpPort string
)

type Response struct {
	Message string `json:"message"`
}

func Setup() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
	}
}

func StartServerHttp() {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	r := mux.NewRouter()

	// generate all the routers
	r = Router(r)

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: r,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("Listening on: ", err)
		}
	}()

	log.Info("Server Started...")
	log.Info("Listening on port: " + httpPort)

	<-done
	log.Info("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server Shutdown Failed:%+v", err)
	}
	log.Info("Server Exited Properly")
}

func Run() {
	Setup()
	StartServerHttp()
}
