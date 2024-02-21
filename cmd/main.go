package main

import (
	"github.com/abiosoft/ishell/v2"
	"github.com/amar-jay/dev_cli/pkg/handlers"
	"os"
)

/**
* The main function of this CLI application is to simplify the Amar Jay's development environment setup.
* This CLI will help to install and configure the following tools:
* - NeoVim
* - Tmux
* - Zsh
* - OhMyZsh
*
* The CLI is modifiable to;
* - backup existing configuration files
* - install and configure other tools
* - setup neovim plugins
 */
func main() {
	shell := ishell.New()

	handler := handlers.New("amar-jay/dev_cli", ".backup")

	shell.AddCmd(&ishell.Cmd{
		Name:    "install",
		Aliases: []string{"i"},
		Func:    handler.CheckInstalled,
		Help:    "Install and configure development environment tools",
	})

	// check environment variables `TMUX_CONFIG_PATH` and `XDG_CONFIG_HOME`
	shell.AddCmd(&ishell.Cmd{
		Name:    "variables",
		Aliases: []string{"vars"},
		Func: func(c *ishell.Context) {
			res := c.MultiChoice([]string{"tmux", "nvim"}, "which of them?")

			switch res {
			case 0:
				c.Println("env of the tmux variable", handler.TmuxConfigPath)
			case 1:
				c.Println("env of the neovim variable", handler.NvimConfigPath)
			default:
				c.Println("Invalid Prompt")
			}
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name:    "welcome",
		Aliases: []string{"w"},
		Func:    handler.Greetings,
		Help:    "Send greetings",
	})

	shell.AddCmd(&ishell.Cmd{
		Name:    "backup",
		Aliases: []string{"back"},
		Func: func(c *ishell.Context) {
			// ask if to backup either tmux or neovim or both
			res := c.Checklist([]string{"tmux", "nvim"}, "Which of them to Backup", []int{0, 1})
			switch {
			case len(res) == 2:
				handler.TmuxBackup(c)
				handler.NvimBackup(c)
			case len(res) == 0:
				c.Println("Did not backup any!! 🙄")
			case res[0] == 0:
				handler.TmuxBackup(c)
			case res[0] == 1:
				handler.NvimBackup(c)
			default:
				c.Println("Invalid Prompt")
			}
		},
		Help: "Backup files,tmux / neovim",
	})
	//shell.Println("Hello to Amar Jay Dev setup CLI")
	if len(os.Args) > 1 && os.Args[1] == "exit" {
		err := shell.Process(os.Args[2:]...)
		if err != nil {
			panic(err)
		}
	} else {
		shell.Run()
		shell.Close()
	}

}
