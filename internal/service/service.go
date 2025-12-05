package service

import (
	"fmt"
	"strings"

	"pkg/morse"
)

const text = "АБВГДЕЁЖЗИЙКЛМНОПРССТУФХЦЧШЩЪЫЬЭЮЯ!1234567890)(,:;%@\\|/"

func Decode(input string) string {
	input = strings.ToUpper(input)

	if len(input) == 0 {
		fmt.Println("input is empty")
		return ""
	}

	if strings.ContainsAny(input, text) {
		return morse.ToMorse(input)
	} else {
		return morse.ToText(input)
	}

}
