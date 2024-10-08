package main

import (
	"fmt"
	_ "image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

var mu sync.Mutex
var counter, errors int

func processFile(filePath string, wg *sync.WaitGroup) {
	defer wg.Done()
	if err := updateFileDate(filePath); err != nil {
		log.Printf("Failed to update file %s: %v", filePath, err)
		errors++
	} else {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func main() {
	currTime := time.Now()
	//Clear the terminal before processing
	clearConsole()

	// Get the current directory
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current directory:", currentPath)

	// Get the file list
	files, err := filepath.Glob(filepath.Join(currentPath, "*"))
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	// Go over files
	for _, file := range files {
		ext := strings.ToLower(filepath.Ext(file))
		if ext == ".jpeg" || ext == ".png" || ext == ".jpg" {
			wg.Add(1)
			go processFile(file, &wg)
		}
	}

	wg.Wait()
	timeTotal := time.Since(currTime)
	fmt.Println("________________________________________")
	fmt.Printf("Processing %d files completed (%v).\n\a", counter, timeTotal)
	if errors != 0 {
		fmt.Printf("\u001B[31mError processing %d files.\u001B[0m\n\a", errors)
	}
	// Wait for user input before closing
	fmt.Println("Press Enter to exit...")
	fmt.Scanln() // Wait for user to press Enter
}

// Date-update function
func updateFileDate(filePath string) error {
	// Open file
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Get EXIF
	x, err := exif.Decode(f)
	if err != nil {
		return err
	}

	// Get shooting date
	date, err := x.DateTime()
	if err != nil {
		return err
	}
	return os.Chtimes(filePath, time.Now(), date)
}
