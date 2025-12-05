package main

import (
	"fmt"

	service "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
	morse "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func main() {
	// Простейший тест
	text := "А"
	fmt.Printf("Текст: %q\n", text)

	m := morse.ToMorse(text)
	fmt.Printf("В морзе: %q\n", m)

	t := morse.ToText(m)
	fmt.Printf("Обратно: %q\n", t)

	if text != t {
		fmt.Println("Ошибка: конвертация не симметрична!")
	} else {
		fmt.Println("Успех: конвертация работает!")
	}

	result, err := service.AutoDetectAndConvert(text)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	if result == m {
		fmt.Println("Успех: автоматическая детекция работает")
	} else {
		fmt.Println("Ошибка: автоматическая детекция не работает")
	}
}
