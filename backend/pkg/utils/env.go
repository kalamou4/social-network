package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// LoadEnvFile reads environment variables from a file and sets them.
// It returns any error encountered during the process.
func LoadEnvFile(filePath string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open env file: %w", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Keep track of the line number for error reporting
	lineNum := 0

	// Scan the file line by line
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Split the line into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("malformed env variable at line %d: %s", lineNum, line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove surrounding quotes if present
		value = strings.Trim(value, `"'`)

		// Set the environment variable
		err := os.Setenv(key, value)
		if err != nil {
			return fmt.Errorf("failed to set env variable %s: %w", key, err)
		}
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading env file: %w", err)
	}

	return nil
}
