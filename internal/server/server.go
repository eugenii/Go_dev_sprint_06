package server

import (
	"fmt"
	"log"
	"net/http"

	handlers "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

func NewServer(port string) *http.Server {
	// Создаем маршрутизатор
	mux := http.NewServeMux()

	// Регистрируем обработчики
	mux.HandleFunc("/", handlers.HandleMain)
	mux.HandleFunc("/convert", handlers.HandleConvert)

	// Настраиваем сервер
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	return server
}

func Start(port string) {
	server := NewServer(port)

	fmt.Printf("Сервер запущен на http://localhost:%s\n", port)
	fmt.Println("Для остановки нажмите Ctrl+C")

	log.Fatal(server.ListenAndServe())
}
