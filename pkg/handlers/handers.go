package handlers

import (
	"os"
	"os/user"
	"path/filepath"
)

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
		Installed:      false,
		RepoName:       repoName,
		NvimConfigPath: filepath.Join(confDir, "nvim"),
		TmuxConfigPath: filepath.Join(usr.HomeDir, ".tmux"),
	}

	// create backup dir if it does not exist
	if _, err := os.Stat(h.BackupDir); err != nil {
		println("Creating backup directory....")
		if err := os.MkdirAll(h.BackupDir, 0755); err != nil {
			panic(err)
		}
	}

	return h
}
