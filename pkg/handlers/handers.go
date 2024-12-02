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
		return fmt.Errorf("error creating symlink for tmux: %v", err)
	}

	// tmux default config file and its symlink
	tmuxSymFile := filepath.Join(b.HomeDir, ".tmux.conf")
	tmuxFile := filepath.Join(b.TmuxConfigPath, "tmux.conf")

	if err := utils.CreateSymlink(tmuxFile, tmuxSymFile); err != nil {
		return fmt.Errorf("error creating symlink for tmux: %v", err)
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

	return h
}

func (h Handlers) Extras() {

	// create backup dir if it does not exist
	if _, err := os.Stat(h.BackupDir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(h.BackupDir, 0755); err != nil {
				panic(fmt.Errorf("failed to create backup directory: %w", err))
			}
		} else {
			panic(fmt.Errorf("failed to check backup directory: %w", err))
		}
	} else {
		println("Backup directory already exists.")
	}
	// check if neovim and tmux config is available if not install
	_, err := os.Stat(h.NvimConfigPath)
	_, err2 := os.Stat(h.TmuxConfigPath)
	println(h.NvimConfigPath)
	println(h.TmuxConfigPath)
	if err != nil || err2 != nil {
		if os.IsNotExist(err) || os.IsNotExist(err2) {
			println("neovim/tmux files dont' exist")
			//h.CloneRepo(CliHandlerContext{})
		}
	}

	// all relevant symbolic links ie. tmux,
	if err := h.symlinks(); err != nil {
		panic(err)
	}

}
