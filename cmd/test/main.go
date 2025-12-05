package main

import (
	"fmt"

	"internal/service"
)

func main() {
	// Тест 1: Текст в Морзе
	fmt.Println("=== Тест 1: Текст 'привет' ===")
	result, err := service.AutoDetectAndConvert("привет")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Результат: %q\n", result)
	}

	// Тест 2: Морзе в текст
	fmt.Println("\n=== Тест 2: Код Морзе ===")
	result2, err := service.AutoDetectAndConvert(".--. .-. .. .-- . -")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Результат: %q\n", result2)
	}

	// Тест 3: С файлами из задания
	fmt.Println("\n=== Тест 3: Содержимое файла test ===")
	// Прочитайте содержимое файла test и передайте в функцию
}
