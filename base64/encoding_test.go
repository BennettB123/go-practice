package main

import (
	"encoding/base64"
	"testing"
)

func TestEncoding(t *testing.T) {
	tests := []struct {
		name string
		text string
	}{
		{"sentence", "encode this for me, please"},
		{"numbers", "1234567890"},
		{"symbols", "!@#$$%^&*()"},
		{"chinese text", "请帮我编码一下"},
		{"empty string", ""},
		{"1 character", "1"},
		{"2 characters", "12"},
		{"3 characters", "123"},
	}

	for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
			expected := base64.StdEncoding.EncodeToString([]byte(test.text))
			actual := encode(test.text)

			if expected != actual {
				t.Errorf("expected [%s] but got [%s]", expected, actual)
			}
        })
    }
}
