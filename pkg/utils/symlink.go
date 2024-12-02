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
//
// CreateSymlink ensures a valid symlink from src to dst
func CreateSymlink(src, dst string) error {
	// Verify if dst is a symlink
	if IsSymlink(dst) {
		target, err := os.Readlink(dst)
		if err != nil {
			return fmt.Errorf("failed to read symlink target: %w", err)
		}
		// If the symlink points to the correct target, return
		if target == src {
			fmt.Printf("Symlink '%v' -> '%v' already exists.\n", dst, src)
			return nil
		}
		// Remove incorrect symlink
		fmt.Printf("Removing incorrect symlink '%v' -> '%v'.\n", dst, target)
		if err := os.Remove(dst); err != nil {
			return fmt.Errorf("failed to remove invalid symlink: %w", err)
		}
	}

	// Remove existing file/directory at dst if it's not a symlink
	if _, err := os.Stat(dst); err == nil {
		//fmt.Printf("Removing existing path '%v'.\n", dst)
		if err := os.RemoveAll(dst); err != nil {
			return fmt.Errorf("failed to remove existing path: %w", err)
		}
	}

	// Create a new symlink
	//fmt.Printf("Creating symlink '%v' -> '%v'.\n", dst, src)
	if err := os.Symlink(src, dst); err != nil {
		if os.IsExist(err) {
			return fmt.Errorf("tmux config doesn't match Amar Jay's config (HINT: DELETE & DOWNLOAD IT)")
		}
		return fmt.Errorf("failed to create symlink: %w", err)
	}

	return nil
}
