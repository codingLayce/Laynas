package model

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func ZipFolder(source string) (string, error) {
	// Temp zip file which will be sent to the client
	f, err := os.CreateTemp(".", "tmp-*.zip")
	if err != nil {
		return "", err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	// GO through all files of the source
	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Header of the current file
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Set compression
		header.Method = zip.Deflate

		// Set relative path of the file
		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			header.Name += "/"
		}

		// Save content of the file
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		// Dir has no content, only header
		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err
	})

	if err != nil {
		return "", err
	}

	return f.Name(), err
}
