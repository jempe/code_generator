package utils

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// FileExists checks if file exists
func FileExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}
	return false
}

// IsDir checks if path is a directory
//
func IsDirectory(file string) bool {
	if stat, err := os.Stat(file); err == nil && stat.IsDir() {
		return true
	}
	return false
}

// CurrentDirectory returns the directory
//
func CurrentDirectory() string {
	directory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return directory
}

// CopyFile copies files
//
func CopyFile(source string, destination string) (err error) {
	if FileExists(source) && !FileExists(destination) {
		sourceFile, err := os.Open(source)
		if err != nil {
			return err
		}
		defer sourceFile.Close()

		destFile, err := os.Create(destination)
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, sourceFile)
		if err != nil {
			return err
		}

		err = destFile.Sync()
		if err != nil {
			return err
		}
	} else {
		err = errors.New("source file doesn't exist or destination file already exists")
	}

	return err
}

// MoveFile move files
//
func MoveFile(source string, destination string) (err error) {
	if FileExists(source) {
		err = os.Rename(source, destination)
		if err != nil {
			return err
		}
	} else {
		err = errors.New("source file " + source + " doesn't exist")
	}

	return err
}

// MakeDirectory creates new directory
//
func MakeDirectory(directory string, permissions os.FileMode) (err error) {
	if !FileExists(directory) {
		err = os.MkdirAll(directory, permissions)
		if err != nil {
			return err
		}
	} else {
		err = errors.New("directory " + directory + " already exists")
	}

	return err
}

//SaveToFile saves content to a file
//
func SaveToFile(file string, content string, permissions os.FileMode) (err error) {
	if !FileExists(file) {
		err = ioutil.WriteFile(file, []byte(content), permissions)
		if err != nil {
			return err
		}
	} else {
		err = errors.New("file " + file + " already exists")
	}

	return err
}

//Remove Last slash
//
func RemoveLastSlash(path string) string {
	return strings.TrimSuffix(path, "/")
}
