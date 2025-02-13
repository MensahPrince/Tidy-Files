package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	fmt.Println("Hello World!")
	sysDir()
}

func sysDir() {
	// Define path to Desktop
	osDir := "C:\\Users\\Mensa\\Desktop"
	fname := "log.txt" // Log file name
	testDest := "C:\\Users\\Mensa\\Desktop\\test"
	// Read directory contents
	files, err := os.ReadDir(osDir)
	if err != nil {
		log.Fatal(err)
	}

	// Create log file
	element, err := os.Create(filepath.Join(osDir, fname))
	if err != nil {
		log.Fatal(err)
	}
	defer element.Close() // Ensure file closes properly

	// Compile regex for .lnk files (Windows shortcuts)
	re, err := regexp.Compile(`(?i)\.lnk$`) // Case-insensitive match for `.lnk`
	if err != nil {
		log.Fatal(err)
	}

	// Loop through files
	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		if re.MatchString(file.Name()) {
			continue // Skip .lnk files
		}else{
			moveDirs(osDir, testDest, files)
		}

		// Write filenames to log.txt
		if _, err := element.WriteString(file.Name() + "\n"); err != nil {
			log.Fatal(err)
		}
	}
}

 
func moveDirs(src, newLoc string, files []os.DirEntry) {
	for _, file := range files {
		err := os.Rename(filepath.Join(src, file.Name()), filepath.Join(newLoc, file.Name()))

		if err != nil {
			log.Println(err)
		}
	}
}