package utils

import (
	"fmt"
	"os"
)

func ReadFile(path string) string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.ReadFile(fmt.Sprintf("%s/mock/%s", dir, path))
	if err != nil {
		panic(err)
	}
	return string(file)
}
