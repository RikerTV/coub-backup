package main

import (
	"os"
)

func DirectorySetup(user string, dir string) (err error) {
	// Creates a directory named user in the dir directory.
	// If the directory already exists, it is not created.
	return CreateDirectory(dir + user)
}

// CreateDirectory creates a directory if it does not exist, and returns an error if it cannot be created and does not already exist.
func CreateDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// FileExists returns true if the file exists, false otherwise.
func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
