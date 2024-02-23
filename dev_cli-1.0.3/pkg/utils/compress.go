package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// GPT generated code
func TarCompress(sourcePath, destinationPath string) error {
	// Create the destination file
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// Create a gzip writer
	gzipWriter := gzip.NewWriter(destinationFile)
	defer gzipWriter.Close()

	// Create a tar writer
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	// Walk through the source file/directory
	err = filepath.Walk(sourcePath, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create a tar header
		header, err := tar.FileInfoHeader(fileInfo, fileInfo.Name())
		if err != nil {
			return err
		}

		// Update the name to relative path
		relPath, err := filepath.Rel(sourcePath, filePath)
		if err != nil {
			return err
		}
		header.Name = relPath

		// Write the header
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		// If it's not a directory, write the file content
		if !fileInfo.Mode().IsDir() {
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(tarWriter, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// GPT generated code
func TarDecompress(sourcePath, destinationPath string) error {
	// Open the source file
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Create a gzip reader
	gzipReader, err := gzip.NewReader(sourceFile)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	// Create a tar reader
	tarReader := tar.NewReader(gzipReader)

	// Iterate through the files in the tar archive
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break // End of tar archive
		}
		if err != nil {
			return err
		}

		// Construct the destination file path
		destPath := filepath.Join(destinationPath, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			// Create directories
			if err := os.MkdirAll(destPath, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			// Create parent directories if not exists
			if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
				return err
			}
			// Create file
			outFile, err := os.Create(destPath)
			if err != nil {
				return err
			}
			defer outFile.Close()

			// Copy file data
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported tar type: %d in file %s", header.Typeflag, header.Name)
		}
	}

	return nil
}
