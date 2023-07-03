package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// LogToFile creates a file with a given suffix and writes the content into it.
func LogToFile(suffix, content string) {
	// Get current time and format it as a string
	timestamp := time.Now().Format("20060102-150405.11.000") // "yyyyMMdd-HHmm.ss.sss"

	// Create a new file with the timestamp and suffix in the name
	file, fileErr := os.Create(fmt.Sprintf("output-logs\\%s-download-%s.txt", timestamp, suffix))
	if fileErr != nil {
		fmt.Println("Failed to create log file:", fileErr)
		return
	}
	defer file.Close()

	// Write the content to the file
	_, writeErr := io.WriteString(file, content)
	if writeErr != nil {
		fmt.Println("Failed to write to log file:", writeErr)
		return
	}

	// Ensure the content is written to disk
	syncErr := file.Sync()
	if syncErr != nil {
		fmt.Println("Failed to write to log file:", syncErr)
		return
	}
}
