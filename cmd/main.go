package main

import (
	"log"
	"os"

	server "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаем логгер
	logger := log.New(os.Stdout, "MORSE: ", log.Ldate|log.Ltime)

	// Создаем сервер
	srv := server.NewServer(logger)

	// Запускаем сервер
	logger.Println("Запуск конвертера азбуки Морзе...")
	logger.Println("Сервер доступен по адресу: http://localhost:8080")

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal("Ошибка сервера: ", err)
	}
}
