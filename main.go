package main

import (
	"fmt"
	_ "image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

func main() {
	// Get current path
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current directory:", currentPath)

	// Get file list
	files, err := filepath.Glob(filepath.Join(currentPath, "*"))
	if err != nil {
		log.Fatal(err)
	}

	// Go through every file in the directory
	for _, file := range files {
		ext := filepath.Ext(file)
		if ext == ".jpeg" || ext == ".jpg" || ext == ".png" {
			fmt.Printf("Processing file: %s\n", file)

			// Open file
			f, err := os.Open(file)
			if err != nil {
				log.Printf("Error opening file %s: %v\n", file, err)
				continue
			}

			// Get EXIFs
			x, err := exif.Decode(f)
			f.Close() // Закрываем файл, так как он больше не нужен
			if err != nil {
				log.Printf("Error decoding EXIF data from file %s: %v\n", file, err)
				continue
			}

			// Get a shot date
			tm, err := x.DateTime()
			if err != nil {
				log.Printf("No DateTime found in EXIF for file %s: %v\n", file, err)
				continue
			}

			// Set the shot date as date of changing file
			err = os.Chtimes(file, time.Now(), tm)
			if err != nil {
				log.Printf("Error setting modification time for file %s: %v\n", file, err)
				continue
			}

			fmt.Printf("Set modification time to %v for file %s\n", tm, file)
		}
	}
}
