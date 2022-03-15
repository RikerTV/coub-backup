package main

import (
	"os"
	"strconv"
	"strings"
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

func CreateCoubDir(dir string, coub Coub) (outdir string, err error) {

	dir = strings.TrimRight(dir, "/")
	CoubYear := coub.CreatedAt.Year()
	CoubMonth := coub.CreatedAt.Month()

	err = CreateDirectory(dir + "/" + strconv.Itoa(CoubYear))
	if err != nil {
		return "", err
	}

	err = CreateDirectory(dir + "/" + strconv.Itoa(CoubYear) + "/" + strconv.Itoa(int(CoubMonth)))
	if err != nil {
		return "", err
	}

	outdir = dir + "/" + strconv.Itoa(CoubYear) + "/" + strconv.Itoa(int(CoubMonth)) + "/" + coub.Title
	err = CreateDirectory(outdir)
	if err != nil {
		return "", err
	}

	return outdir, nil
}
