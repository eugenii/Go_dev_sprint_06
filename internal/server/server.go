package server

import (
	"log"
	"net/http"
	"time"

	handlers "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

func NewServer(logger *log.Logger) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HandleMain)

	// Второй хендлер (новый)
	mux.HandleFunc("/upload", handlers.HandleUpload)

	return &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}
