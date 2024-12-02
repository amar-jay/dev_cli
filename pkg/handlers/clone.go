package handlers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// ignore unused
func ghForkRepo(c HandlerContext, repoName, tempDir string) error {
	out, err := exec.Command("gh", "repo", "view", repoName, "--json", "fork", "--jq", ".fork").Output()
	if err != nil {
		return err
	}
	forked := strings.TrimSpace(string(out))
	// check if user has forked repo, use that else fork it
	if forked == "true" {
		c.Println("You have already forked the repository. Cloning your fork.")
		out, err := exec.Command("gh", "repo", "view", repoName, "--json", "html_url", "--jq", ".parent.fork.html_url").Output()
		if err != nil {
			return fmt.Errorf("error cloning repo:%v", err)
		}

		parentURL := strings.TrimSpace(string(out))

		cmd := exec.Command("git", "clone", parentURL, tempDir)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error cloning the repo: %v", err)
		}
	} else {

		c.Println("Forking the repository...")
		cmd := exec.Command("gh", "repo", "fork", repoName, "--clone")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error forking the repo: %v", err)
		}

		// Extracting the repository name from the full path
		repoDir := repoName[strings.LastIndex(repoName, "/")+1:]

		err := os.Rename(repoDir, tempDir)
		if err != nil {
			return fmt.Errorf("error renaming the repo : %v", err)
		}

	}
	println("forked and clone repo successfully ...")
	return nil
}

func gitForkRepo(_ HandlerContext, repoName, tempDir string) error {
	println("git", "clone", "--recurse-submodules", "https://github.com/"+repoName+".git", tempDir)
	cmd := exec.Command("git", "clone", "--recurse-submodules", "https://github.com/"+repoName+".git", tempDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error cloning the repo: %v", err)
	}
	return nil
}

// clone github config to path and load
func (b *Handlers) CloneRepo(c HandlerContext) {
	tempDir, err := os.MkdirTemp("", "dev_cli-*")
	if err != nil {
		c.Println("error creating temporary directory:", err.Error())
		return
	}
	defer os.RemoveAll(tempDir)

	//NOTE: Due to some complications with gh, using git by default
	print("NOTE: Due to some complications with gh, using git by default")
	c.Println("gh CLI is not installed. Cloning the", b.RepoName, "repository using git to", tempDir)
	if err := gitForkRepo(c, b.RepoName, tempDir); err != nil {
		c.Println("git CLI clone error: ", err.Error())
		return
	}
	/*
		    // Check if the gh CLI is installed
				if _, err := exec.LookPath("gh"); err != nil {
					c.Println("gh CLI is not installed. Cloning the", b.RepoName, "repository using git.")
					gitForkRepo(c, b.RepoName, tempDir)
				}else {
						  println("Using gh cli", fi)
						if err := ghForkRepo(c, b.RepoName, tempDir); err != nil {
							c.Println("GH fork error: ", err)
							return
						}
			  }
	*/

	configs := map[string]string{"nvim": b.NvimConfigPath, "tmux": b.TmuxConfigPath}
	for name, newpath := range configs {
		oldpath := filepath.Join(tempDir, name)
		if _, err := os.Lstat(newpath); os.IsNotExist(err) {
			if err := os.Rename(oldpath, newpath); err != nil {
				c.Println("1, error copying", oldpath, "to", newpath, err.Error())
				return
			}
		}

		if err = os.RemoveAll(newpath); err != nil {
			c.Println("error removing previous config", err.Error())
			return
		}

		if err := os.Rename(oldpath, newpath); err != nil {
			
			if _, err = os.Stat(oldpath); err != nil && os.IsNotExist(err){
				println(">>" + oldpath)
				return
			}
			
			if _, err = os.Stat(newpath); err != nil  && os.IsNotExist(err) {
				println(">>" + newpath)
				return
			}

			if err != nil {
				c.Println("2, error copying", oldpath, "to", newpath, err.Error())
				return
			}
		}
	}

	c.Println("Copied config successfully...")
}
