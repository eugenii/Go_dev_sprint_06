// test_handler.go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	handlers "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

func main() {
	// Создаем тестовый запрос
	req := httptest.NewRequest("GET", "/", nil)

	// Создаем ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()

	// Вызываем обработчик
	handlers.HandleMain(rr, req)

	// Проверяем статус код
	if rr.Code != http.StatusOK {
		fmt.Printf("Ошибка: статус код %d\n", rr.Code)
	}

	// Проверяем Content-Type
	contentType := rr.Header().Get("Content-Type")
	if contentType != "text/html; charset=utf-8" {
		fmt.Printf("Ошибка Content-Type: %s\n", contentType)
	} else {
		fmt.Println("Content-Type правильный")
	}

	// Проверяем, что ответ не пустой
	body := rr.Body.String()
	if len(body) == 0 {
		fmt.Println("Ошибка: пустое тело ответа")
	} else {
		fmt.Printf("Ответ содержит %d байт\n", len(body))

		// Проверяем, что это HTML
		if strings.Contains(strings.ToLower(body), "<html") {
			fmt.Println("Ответ содержит HTML")
		}
		if strings.Contains(strings.ToLower(body), "<form") {
			fmt.Println("Ответ содержит форму")
		}

		// Можно вывести первые 200 символов для проверки
		if len(body) > 200 {
			fmt.Printf("Первые 200 символов:\n%s...\n", body[:500])
		} else {
			fmt.Printf("Содержимое:\n%s\n", body)
		}
	}
}
