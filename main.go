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
	// Получаем текущую директорию
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current directory:", currentPath)

	// Получаем список файлов в папке
	files, err := filepath.Glob(filepath.Join(currentPath, "*"))
	if err != nil {
		log.Fatal(err)
	}

	// Проходим по каждому файлу в папке
	for _, file := range files {
		ext := filepath.Ext(file)
		if ext == ".jpeg" || ext == ".jpg" || ext == ".png" {
			fmt.Printf("Processing file: %s\n", file)

			// Открываем файл
			f, err := os.Open(file)
			if err != nil {
				log.Printf("Error opening file %s: %v\n", file, err)
				continue
			}

			// Получаем метаданные EXIF
			x, err := exif.Decode(f)
			f.Close() // Закрываем файл, так как он больше не нужен
			if err != nil {
				log.Printf("Error decoding EXIF data from file %s: %v\n", file, err)
				continue
			}

			// Извлекаем дату съёмки
			tm, err := x.DateTime()
			if err != nil {
				log.Printf("No DateTime found in EXIF for file %s: %v\n", file, err)
				continue
			}

			// Устанавливаем дату изменения файла на дату съёмки
			err = os.Chtimes(file, time.Now(), tm)
			if err != nil {
				log.Printf("Error setting modification time for file %s: %v\n", file, err)
				continue
			}

			fmt.Printf("Set modification time to %v for file %s\n", tm, file)
		}
	}
}
