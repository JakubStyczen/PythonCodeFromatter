package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatLine(t *testing.T) {
	tests := []struct {
		input_string, expected string
	}{
		{"Kox = 2", "Kox = 2"},
		{"Kox  =  2", "Kox = 2"},
		{" Kox=2 ", "Kox = 2"},
	}

	for _, tt := range tests {
		t.Run("Correct Fromatting", func(t *testing.T) {
			result, _ := FormatVariableDeclaration(tt.input_string)
			assert.Equal(t, result, tt.expected)
		})
	}

}
