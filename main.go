package main

import (
	"fmt"
	_ "image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

func processFile(filePath string, wg *sync.WaitGroup) {
	defer wg.Done()
	if err := updateFileDate(filePath); err != nil {
		log.Printf("Failed to update file %s: %v", filePath, err)
	}
}

func main() {
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
		ext := filepath.Ext(file)
		if ext == ".jpeg" || ext == ".png" || ext == ".jpg" {
			wg.Add(1)
			go processFile(file, &wg)
		}
	}

	wg.Wait()
	fmt.Println("Processing completed.")
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
