package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func appendFiles(dirname string, zipw *zip.Writer) error {
	var files []string
	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {

		} else {
			files = append(files, path)

		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, filename := range files {

		file, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("Failed to open %s: %s", filename, err)
		}
		defer file.Close()
		wr, err := zipw.Create(filename)
		if err != nil {
			msg := "Failed to create entry for %s in zip file: %s"
			return fmt.Errorf(msg, filename, err)
		}

		if _, err := io.Copy(wr, file); err != nil {
			return fmt.Errorf("Failed to write %s to zip: %s", filename, err)
		}
	}
	return nil
}

func main() {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile("test.zip", flags, 0644)
	if err != nil {
		log.Fatalf("Failed to open zip for writing: %s", err)
	}
	defer file.Close()

	path := "C:\\Users\\name\\Documents\\New folder"

	zipw := zip.NewWriter(file)
	defer zipw.Close()

	if err := appendFiles(path, zipw); err != nil {
		log.Fatalf("Failed to add file %s to zip: %s", path, err)
	}

}
