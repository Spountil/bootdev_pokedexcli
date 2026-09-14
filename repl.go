package main

import "strings"

func cleanInput(text string) []string {
	var result []string
	words := strings.Fields(text)

	for _, word := range words {
		result = append(result, strings.ToLower(word))
	}
	return result
}
