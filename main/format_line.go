package main

import (
	"errors"
	"strings"
)

func FilterSpacesInStringSlice(slice []string) []string {
	var result []string
	for _, elem := range slice {
		if elem != "" {
			result = append(result, elem)
		}
	}
	return result
}

func FormatVariableDeclaration(line string) (string, error) {
	// Line should look like this: varName = value
	words := strings.Split(line, " ")
	filteredWords := FilterSpacesInStringSlice(words)
	if len(filteredWords) != 3 {
		return line, errors.New("wrong data format")
	}
	return strings.Join(filteredWords, " "), nil
}
