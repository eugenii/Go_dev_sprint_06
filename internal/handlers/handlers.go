package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	service "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

var ProjectRoot string

func SetProjectRoot(root string) {
	ProjectRoot = root
}

func HandleMain(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	htmlPath := ProjectRoot + "./index.html"
	file, err := os.Open(htmlPath)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.Copy(w, file)
}

func HandleConvert(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Базовые проверки
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "No file uploaded", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Читаем
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Read error", http.StatusInternalServerError)
		return
	}

	// Конвертируем
	original := string(data)
	converted, err := service.AutoDetectAndConvert(original)
	if err != nil {
		http.Error(w, "Conversion failed", http.StatusInternalServerError)
		return
	}

	// Формируем результат
	result := fmt.Sprintf("оригинал:\n%s\n\nконвертированный:\n%s", original, converted)

	// Сохраняем с тем же именем
	filename := header.Filename
	if err := os.WriteFile(filename, []byte(result), 0644); err != nil {
		http.Error(w, "Save failed", http.StatusInternalServerError)
		return
	}

	// Ответ
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("OK - файл сохранен: " + filename))
}
