package utils

import (
	"fmt"
	"io/fs"
	"os"
)

func ReadDir(dirPath string) ([]fs.DirEntry, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %s", err)
	}

	return entries, nil
}
