package main

import (
	"fmt"
	"log"
	"os"

	"github.com/amar-jay/dev_cli/pkg/handlers"
	"github.com/abiosoft/ishell"
	"github.com/urfave/cli/v2"
)

/**
* The shell function of this CLI application is to simplify the Amar Jay's development environment setup.
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
func Shell(handler handlers.Handlers, osArgs []string) {
	handler.Extras()
	shell := ishell.New()

	// Wrapper to adapt *ishell.Context to HandlerContext
	wrapShellHandler := func(f func(handlers.HandlerContext)) func(c *ishell.Context) {
		return func(c *ishell.Context) {
			f(c) // Pass the *ishell.Context directly, as it implements HandlerContext
		}
	}

	shell.AddCmd(&ishell.Cmd{
		Name:    "download",
		Aliases: []string{"i"},
		Func:    wrapShellHandler(handler.CloneRepo),
		Help:    "Downloading of amar-jay's setup",
	})

	shell.AddCmd(&ishell.Cmd{
		Name:    "install",
		Aliases: []string{"i"},
		Func:    wrapShellHandler(handler.CheckInstalled),
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
		Help: "Check for present environment variables",
	})

	shell.AddCmd(&ishell.Cmd{
		Name:    "welcome",
		Aliases: []string{"w"},
		Func:    wrapShellHandler(handler.Greetings),
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
	if len(osArgs) > 1 && osArgs[1] == "exit" {
		err := shell.Process(osArgs[2:]...)
		if err != nil {
			panic(err)
		}
	} else {
		shell.Run()
		shell.Close()
	}

}

/**
* The cli function of this CLI application is to simplify Amar Jay's development environment setup.
* This CLI will help to install and configure the following tools:
* - NeoVim
* - Tmux
* - Zsh
* - OhMyZsh
*
* The CLI is modifiable to:
* - backup existing configuration files
* - install and configure other tools
* - setup neovim plugins
 */
func Cli(handler handlers.Handlers, osArgs []string) {

	wrapCliHandler := func(f func(handlers.HandlerContext)) cli.ActionFunc {
		return func(c *cli.Context) error {
			f(handlers.HandlerContext(handlers.CliHandlerContext{Context: c})) // Pass the *ishell.Context directly, as it implements HandlerContext
			return nil
		}
	}
	app := &cli.App{
		Name:  "devsetup",
		Usage: "Simplify Amar Jay's development environment setup",
		Commands: []*cli.Command{
			{
				Name:    "download",
				Aliases: []string{"d", "down"},
				Usage:   "Download Amar Jay's setup",
				Action:  wrapCliHandler(handler.CloneRepo),
			},
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "Install and configure development environment tools",
				Action:  wrapCliHandler(handler.CheckInstalled),
			},
			{
				Name:    "variables",
				Aliases: []string{"vars"},
				Usage:   "Check for present environment variables",
				Action: func(c *cli.Context) error {
					choices := []string{"tmux", "nvim"}
					log.Println("Select an option:")
					for i, choice := range choices {
						log.Printf("[%d] %s\n", i, choice)
					}

					var choice int
					_, err := fmt.Scan(&choice)
					if err != nil || choice < 0 || choice >= len(choices) {
						log.Println("Invalid choice")
						return nil
					}

					switch choice {
					case 0:
						log.Println("Env of the tmux variable:", handler.TmuxConfigPath)
					case 1:
						log.Println("Env of the neovim variable:", handler.NvimConfigPath)
					}
					return nil
				},
			},
			{
				Name:    "welcome",
				Aliases: []string{"w", "hello"},
				Usage:   "Send greetings",
				Action:  wrapCliHandler(handler.Greetings),
			},
			{
				Name:    "shell",
				Aliases: []string{"sh"},
				Usage:   "Send greetings",
				Action: func(c *cli.Context) error {
					log.Println("Running shell...")
					fmt.Println(c.Args().Slice())
					Shell(handler, c.Args().Slice())
					return nil
				},
			},
			{
				Name:    "backup",
				Aliases: []string{"back"},
				Usage:   "Backup files, tmux/neovim",
				Action: func(c *cli.Context) error {
					log.Println("Which of them to backup?")
					choices := []string{"tmux", "nvim"}
					for i, choice := range choices {
						log.Printf("[%d] %s\n", i, choice)
					}

					var selected []int
					_, err := fmt.Scan(&selected)
					if err != nil {
						log.Println("Invalid input")
						return nil
					}

					if len(selected) == 0 {
						log.Println("Did not backup any! 🙄")
					} else {
						for _, s := range selected {
							switch s {
							case 0:
								wrapCliHandler(handler.TmuxBackup)(c)
							case 1:
								wrapCliHandler(handler.NvimBackup)(c)
							}
						}
					}
					return nil
				},
			},
		},
	}

	err := app.Run(osArgs)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	handler := handlers.New("amar-jay/dev_cli", ".dev_cli")
	Cli(handler, os.Args)
}
