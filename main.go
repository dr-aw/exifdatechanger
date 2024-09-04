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

var mu sync.Mutex

func processFile(filePath string, count *int, wg *sync.WaitGroup) {
	defer wg.Done()
	if err := updateFileDate(filePath, count); err != nil {
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
	var count int
	// Go over files
	for _, file := range files {
		ext := filepath.Ext(file)
		if ext == ".jpeg" || ext == ".png" || ext == ".jpg" {
			wg.Add(1)
			go processFile(file, &count, &wg)
		}
	}

	wg.Wait()
	fmt.Printf("Processing %d files completed.\n\a", count)
}

// Date-update function
func updateFileDate(filePath string, count *int) error {
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
	mu.Lock()
	*count++
	mu.Unlock()
	return os.Chtimes(filePath, time.Now(), date)
}
