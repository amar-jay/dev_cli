package handlers

import (
	"github.com/gookit/color"
)

func (b *Handlers) Greetings(c HandlerContext) {
	c.Println("")
	color.Style{color.FgBlue.Darken(), color.OpBold}.Printf("%-65s\n", "==================================================================")
	color.Style{color.FgBlue.Darken(), color.OpBold}.Printf("%-65s\n", "          Welcome to amarjay's Dev setup!")
	color.Style{color.FgBlue.Darken(), color.OpBold}.Printf("%-65s\n", "==================================================================")
	c.Println("")
	c.Println("This is amarjay's nvim-config, built with Neovim, LazyVim, Mason and TMUX")
	c.Println("")
	c.Println("Neovim is a powerful text editor that provides an intuitive and extensible")
	c.Println("editing environment. With LazyVim, we bring laziness to the next level,")
	c.Println("automating repetitive tasks and reducing cognitive load. Mason enhances")
	c.Println("Neovim's capabilities, allowing seamless integration of various tools")
	c.Println("and workflows.")
	c.Println("")
	c.Println("With this configuration, you'll unleash your productivity, effortlessly")
	c.Println("navigating through code, writing prose, or tinkering with configurations.")
	c.Println("Say goodbye to mundane tasks and hello to a world of efficient editing!")
	c.Println("")
	c.Println("In the case where you have problems with emojis and icons,")
	c.Println("download and install any nerd fonts on your PC.")
	// TODO: fetch link to nerdfonts
	c.Println("https://github.com/" + b.RepoName)
	c.Println("")
	color.Style{color.FgGreen, color.OpItalic}.Println("Setup completed successfully.")
	c.Println("")
	color.Style{color.FgBlue, color.OpBold}.Println("Happy Hacking!!😉")
	c.Println("")
	color.Style{color.FgBlue.Darken(), color.OpBold}.Printf("%-65s\n", "==================================================================")
	c.Println("")
}
