package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// Init инициализирует обработчики с логгером

// func InitHandlers(l *log.Logger) {
// 	logger = l
// }

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
	_, err = io.Copy(w, file)
	if err != nil {
		// Если ошибка при копировании
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}
}

// HandleUpload обрабатывает загрузку файла
func HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Парсим форму
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Cannot parse form", http.StatusBadRequest)
		return
	}

	// Получаем файл
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Cannot get file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Читаем файл
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	// Передаем данные в service для конвертации
	content := string(data)
	converted, err := service.AutoDetectAndConvert(content)
	if err != nil {
		http.Error(w, "Conversion error: "+err.Error(),
			http.StatusInternalServerError)
		return
	}

	fmt.Printf("Конвертация: %d -> %d символов\n",
		len(content), len(converted))

	// Создаем локальный файл для результата
	// Генерируем уникальное имя файла
	timestamp := time.Now().UTC().Format("20060102_150405")
	originalExt := filepath.Ext(header.Filename)
	baseName := header.Filename[:len(header.Filename)-len(originalExt)]

	outputFilename := fmt.Sprintf("%s_converted_%s.txt", baseName, timestamp)

	// Записываем результат в файл
	err = os.WriteFile(outputFilename, []byte(converted), 0644)
	if err != nil {
		http.Error(w, "Cannot save file: "+err.Error(),
			http.StatusInternalServerError)
		return
	}

	fmt.Printf("Результат сохранен в: %s\n", outputFilename)

	// Возвращаем результат
	response := fmt.Sprintf(
		"Конвертация завершена!\n"+
			"Исходный файл: %s\n"+
			"Результат сохранен в: %s\n\n"+
			"Результат конвертации:\n%s",
		header.Filename,
		outputFilename,
		converted,
	)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(response))
}
