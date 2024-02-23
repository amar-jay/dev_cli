package utils

import (
	"fmt"
	"os"
)

// checks wheather a path is a is a symbolic link or not
func IsSymlink(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.Mode().IsRegular()
}

// Creates a new symlink from src to dst.
// If src points to a valid file/directory then it removes it before creating the symlink.
//
//	@param pathLink : is the path of the symbolic link
//	@param path: is the path of the file or directory
func CreateSymlink(src, dst string) error {
	// Check if src is a symlink
	if IsSymlink(dst) {
		//DEBUG: fmt.Printf("Source '%v' is already a symlink.\n", dst)
		return nil
	}

	// Remove existing files/directories at src location
	_, err := os.Stat(dst)
	if err == nil || (!os.IsNotExist(err)) {
		fmt.Println("Removing old source")
		err = os.RemoveAll(dst) //TODO: does this work on all platforms?
		if err != nil {
			return err
		}
	}

	// Create a new symlink
	err = os.Symlink(src, dst)
	if err != nil {
		return err
	}

	return nil
}
