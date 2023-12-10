package utils

import (
	"io"
	"os"
	"path/filepath"
)

func ReadFile(name string) string {
	currentDir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	fullpath := filepath.Join(currentDir, name)

	file, err := os.Open(fullpath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	content, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	return string(content)
}

func WriteFile(filePath string, content []byte) {
	file, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = file.Write(content)

	if err != nil {
		panic(err)
	}
}
