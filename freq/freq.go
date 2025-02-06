package main

import (
	"strings"
)

func wordFrequency(text string) map[string]int {
	if text == ""{
		return make(map[string]int) // возвращаем пустую мапу, можно еще так map[string]int{}
	} 

	freqWord := make(map[string]int, 0)

	buff := ""

	for _, ch := range text{
		if allowedChar(ch) {
			buff += string(ch) // инкримент к числу
			continue
		}

		if buff != "" {
			freqWord[strings.ToLower(buff)]++
			buff = ""
		}
	}

	if buff != "" {
		freqWord[strings.ToLower(buff)]++
	}
	
	return freqWord
}

func allowedChar(el rune) bool {
	// element in 0-9 or A-Z or a-z
	return (el >= 48 && el <= 57) || (el >= 65 && el <= 90) || (el >= 97 && el <= 122) // unicode.IsLetter(el) || unicode.IsDigit(el)
}

// !((входит в диапоз 1) ИЛИ (входит в диапоз 2) ИЛИ (входит в диапоз 3))