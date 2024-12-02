package handlers

import (
	"fmt"
	"os"
	"os/exec"
)

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func execCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute command %s: %v", command, err)
	}
	return nil
}

// NOTE: add progress bar
func (b *Handlers) CheckInstalled(c HandlerContext) {

	// Check if neovim is already installed
	if !commandExists("nvim") {
		c.Println("Neovim is not installed. Installing...")
		// Check the package manager and install neovim accordingly
		if commandExists("apt-get") {
			execCommand("sudo", "apt-get", "install", "neovim", "-y")
		} else if commandExists("pacman") {
			execCommand("sudo", "pacman", "-S", "neovim", "--noconfirm")
		} else if commandExists("brew") {
			execCommand("brew", "install", "neovim")
		} else {
			c.Println("Unable to determine package manager. Please install Neovim manually.")
			return
		}
	} else {
		c.Println("Neovim is already installed.")
	}

	// Check if tmux is already installed
	if !commandExists("tmux") {
		c.Println("Tmux is not installed. Installing...")
		// Check the package manager and install tmux accordingly
		if commandExists("apt-get") {
			execCommand("sudo", "apt-get", "install", "tmux", "-y")
		} else if commandExists("pacman") {
			execCommand("sudo", "pacman", "-S", "tmux", "--noconfirm")
		} else if commandExists("brew") {
			execCommand("brew", "install", "tmux")
		} else {
			c.Println("Unable to determine package manager. Please install Tmux manually.")
			return
		}
	} else {
		c.Println("Tmux is already installed.")
	}

	b.Installed = true
}
