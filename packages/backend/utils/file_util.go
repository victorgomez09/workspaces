package utils

import (
	"fmt"
	"io"
	"os"
)

func MoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("failed removing original file: %s", err)
	}

	return nil
}

func CopyFile(sourcePath string, destPath string) error {
	original, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("failed removing original file: %s", err)
	}
	defer original.Close()

	// Create new file
	new, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed removing original file: %s", err)
	}
	defer new.Close()

	//This will copy
	bytesWritten, err := io.Copy(new, original)
	if err != nil {
		return fmt.Errorf("failed removing original file: %s", err)
	}
	fmt.Printf("Bytes Written: %d\n", bytesWritten)

	return nil
}
