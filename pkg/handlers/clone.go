package handlers

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/abiosoft/ishell/v2"
)

func ghForkRepo(c *ishell.Context, repoName, tempDir string) {
	out, err := exec.Command("gh", "repo", "view", repoName, "--json", "fork", "--jq", ".fork").Output()
	if err != nil {
		c.Println("Error ", err)
		return
	}
	forked := strings.TrimSpace(string(out))
	// check if user has forked repo, use that else fork it
	if forked == "true" {
		c.Println("You have already forked the repository. Cloning your fork.")
		out, err := exec.Command("gh", "repo", "view", repoName, "--json", "html_url", "--jq", ".parent.fork.html_url").Output()
		if err != nil {
			c.Println("Error cloning repo", err)
			return
		}

		parentURL := strings.TrimSpace(string(out))

		cmd := exec.Command("git", "clone", parentURL, tempDir)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			c.Println("Error cloning the repo", err)
			return
		}
	} else {

		c.Println("Forking the repository...")
		cmd := exec.Command("gh", "repo", "fork", repoName, "--clone")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			c.Println("Error forking the repo", err)
			return
		}

		repoDir := repoName[strings.LastIndex(repoName, "/")+1:] // Extracting the repository name from the full path
		err := os.Rename(repoDir, tempDir)
		if err != nil {
			c.Println("Error renaming the repo", err)
			return
		}

	}
	c.Println("forked and clone repo successfully ...")
}

func gitForkRepo(c *ishell.Context, repoName, tempDir string) {
	cmd := exec.Command("git", "clone", "https://github.com/"+repoName+".git", tempDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		c.Println("Error cloning the repo", err)
		return
	}
}

// clone github config to path and load
func (b *Handlers) CloneRepo(c *ishell.Context) {
	tempDir, err := os.MkdirTemp("", "dev_cli*")
	if err != nil {
		c.Println("Error creating temporary directory:", err)
		return
	}
	defer os.RemoveAll(tempDir)

	// Check if the gh CLI is installed
	if _, err := exec.LookPath("gh"); err == nil {
		ghForkRepo(c, b.RepoName, tempDir)
	} else {
		c.Println("gh CLI is not installed. Cloning the", b.RepoName, "repository using git.")
		gitForkRepo(c, b.RepoName, tempDir)
	}

	configs := map[string]string{"nvim": b.NvimConfigPath, "tmux": b.TmuxConfigPath}
	for name, newpath := range configs {
		oldpath := filepath.Join(tempDir, name)
		if err := os.Rename(oldpath, newpath); err != nil {
			c.Println("error copying", oldpath, "to", newpath)
			return
		}
	}

	c.Println("Copied config successfully...")
}
