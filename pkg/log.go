package pkg

import "os"

// OpenLogFile opens or creates a log file at the specified path and returns a pointer to the opened file.
// If the file doesn't exist, it will be created. If it exists, new content will be appended to the file.
// The file is opened in write-only mode with write permissions for owner and group (0644).
// Returns the opened file and any error encountered during the process.
func OpenLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
