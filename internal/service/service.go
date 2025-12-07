package service

import (
	"strings"

	morse "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoDetectAndConvert(data string) (string, error) {
	// Проверяем, является ли строка кодом Морзе
	isMorse := isMorseCode(data)

	if isMorse {
		// Конвертируем Морзе в текст
		text := morse.ToText(data)
		return text, nil
	} else {
		// Конвертируем текст в Морзе
		// Приводим к верхнему регистру для работы с пакетом morse
		upperText := strings.ToUpper(strings.TrimSpace(data))
		morseCode := morse.ToMorse(upperText)
		return morseCode, nil
	}
}

func isMorseCode(data string) bool {
	trimmed := strings.TrimSpace(data)
	if trimmed == "" {
		return false
	}

	// Проверяем, содержит ли строка только символы Морзе
	// Допустимые символы: точка, тире, пробел, табуляция, переносы строк
	for _, r := range trimmed {
		if !(r == '.' || r == '-' || r == ' ' || r == '\t' || r == '\n' || r == '\r') {
			return false
		}
	}

	// Дополнительная проверка: если есть хотя бы один символ Морзе
	// (точка или тире), считаем что это код Морзе
	for _, r := range trimmed {
		if r == '.' || r == '-' {
			return true
		}
	}

	// Только пробелы - не Морзе
	return false
}
