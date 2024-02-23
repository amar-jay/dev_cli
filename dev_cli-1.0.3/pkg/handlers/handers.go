package handlers

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/amar-jay/dev_cli/pkg/utils"
)

/*
symbolic links of the configuration file/dirs from
  - `~/.config/tmux to `~/.tmux`
  - `~/.config/tmux/tmux.conf` to `~/.tmux.conf`
*/
func (b *Handlers) symlinks() error {
	// path to all tmux plugins and other stuff
	tmuxSymDir := filepath.Join(b.HomeDir, ".tmux")

	if err := utils.CreateSymlink(b.TmuxConfigPath, tmuxSymDir); err != nil {
		return fmt.Errorf("Error creating symlink for tmux: %v", err)
	}

	// tmux default config file and its symlink
	tmuxSymFile := filepath.Join(b.HomeDir, ".tmux.conf")
	tmuxFile := filepath.Join(b.TmuxConfigPath, "tmux.conf")

	if err := utils.CreateSymlink(tmuxFile, tmuxSymFile); err != nil {
		return fmt.Errorf("Error creating symlink for tmux: %v", err)
	}
	return nil

}

// holds the configuration paths of the tools.
// Creates a backup directory if it does not exist
func New(repoName, backupDir string) Handlers {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	confDir := os.Getenv("XDG_CONFIG_HOME")
	if confDir == "" {
		confDir = filepath.Join(usr.HomeDir, ".config")
	}

	h := Handlers{
		BackupDir:      filepath.Join(usr.HomeDir, backupDir),
		HomeDir:        usr.HomeDir,
		Installed:      false,
		RepoName:       repoName,
		NvimConfigPath: filepath.Join(confDir, "nvim"),
		TmuxConfigPath: filepath.Join(confDir, "tmux"),
	}

	// create backup dir if it does not exist
	if _, err := os.Stat(h.BackupDir); err != nil {
		println("Creating backup directory....")
		if err := os.MkdirAll(h.BackupDir, 0755); err != nil {
			panic(err)
		}
	}

	// all relevant symbolic links ie. tmux,
	if err := h.symlinks(); err != nil {
		panic(err)
	}

	return h
}
