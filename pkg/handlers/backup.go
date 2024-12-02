package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/amar-jay/dev_cli/pkg/utils"
)

type Handlers struct {
	TmuxConfigPath string // TMUX_CONFIG_PATH is the default path for tmux configuration file
	BackupDir      string // backup directory for configuration files
	NvimConfigPath string // XDG_CONFIG_HOME is the default path for nvim configuration file
	RepoName       string // this is the github repository name, it uses this if it is a fork
	Installed      bool   // tells if whether both neovim and tmux are installed.
	HomeDir        string // path to home directory
}

// TmuxBackup backs up the existing tmux configuration file be it a file or directory
func (b *Handlers) TmuxBackup(c HandlerContext) {
	c.Println("Backing up previous tmux config....")
	currentTime := time.Now().Format("2006-01-02")

	// filename = tmux_20-Feb-21.conf.bak
	backup_filename := fmt.Sprintf("tmux_%s.conf.bak", currentTime)
	backup_filename = filepath.Join(b.BackupDir, backup_filename)

	// if backup file exists, remove it, then create a new one
	if _, err := os.Stat(backup_filename); err == nil {
		c.Println("A backup file exists with same name")
		os.Remove(backup_filename)
	}

	// tar compression of tmux config file
	err := utils.TarCompress(b.TmuxConfigPath, backup_filename)
	if err != nil {
		c.Println("Error creating TMUX backup file", err)
		return
	}

	c.Println("Backup successful")

}

// NvimBackup backs up the existing nvim configuration be it a file or directory
func (b *Handlers) NvimBackup(c HandlerContext) {
	c.Println("Backing up previous nvim config....")
	currentTime := time.Now().Format("2006-01-02")

	// filename = tmux_20-Feb-21.conf.bak
	backup_filename := fmt.Sprintf("nvim_%s.conf.bak", currentTime)
	backup_filename = filepath.Join(b.BackupDir, backup_filename)

	// if backup file exists, remove it, then create a new one
	if _, err := os.Stat(backup_filename); err == nil {
		c.Println("A backup file exists with same name. Remove it first")
		return
		//os.Remove(backup_filename)
	}

	// tar compression of tmux config file
	err := utils.TarCompress(b.NvimConfigPath, backup_filename)
	if err != nil {
		c.Println("Error creating NVIM backup file", err)
		return
	}

	c.Println("Backup successful")

}

// restores the previous tmux configuration backup
func (b *Handlers) TmuxRestore(c HandlerContext) {
	// TODO: based on the backup file name, determone which is the latest Backup file
	// TODO: then untar the file and restore the configuration
	// TODO: if the backup file does not exist, then print a message
	// TODO: in the case of tmux, rerun plugins installation
}

// restores the previous neovim configuration backup
func (b *Handlers) NvimRestore(c HandlerContext) {
	// TODO: based on the backup file name, determone which is the latest Backup file
	// TODO: then untar the file and restore the configuration
	// TODO: if the backup file does not exist, then print a message
}
