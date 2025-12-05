package main

import (
	"flag"
	"fmt"

	server "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Парсим аргументы командной строки
	port := flag.String("port", "8080", "Порт для запуска сервера")
	flag.Parse()

	fmt.Println("=== Конвертер азбуки Морзе ===")
	fmt.Println("Автоматически определяет и конвертирует:")
	fmt.Println("- Текст → код Морзе")
	fmt.Println("- Код Морзе → текст")
	fmt.Println()

	// Запускаем сервер
	server.Start(*port)
}
