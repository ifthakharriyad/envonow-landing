package main

import (
	"log"
	"os"
	"regexp"
	"strings"
)

var logger *log.Logger

// Dirfiles returns all the files in a directory with a specific pattern
func dirfiles(dir string, pat string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		danger(err)
		return nil, err
	}
	files := []string{}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		matched, err := regexp.Match(pat, []byte(entry.Name()))
		if err != nil {
			danger(err)
			continue
		}
		if matched {
			files = append(files, entry.Name())
		}

	}
	return files, nil
}

// Gets full file paths of the files
func filepaths(dir string, files []string) []string {
	for i, file := range files {
		files[i] = strings.Join([]string{dir, file}, "/")
	}
	return files
}

// for logging informations in the program
func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

// for logging errors in the program
func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}
