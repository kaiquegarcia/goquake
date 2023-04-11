package log2memory

import (
	"goquake/internal/config"
	"os"
)

func ReadFileContents() (string, error) {
	data, err := os.ReadFile(config.Get().Filepath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
