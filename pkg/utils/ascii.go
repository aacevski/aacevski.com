package utils

import (
	"os"
	"strings"
)

func LoadASCIIArt(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}
