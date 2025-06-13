package app

import (
	"pswdmng/internal/commands"

	"github.com/spf13/cobra"
)

type App struct {
	rootcmd *cobra.Command
}

func New() *App {
	commands.InitCommands()
	return &App{
		rootcmd: &cobra.Command{},
	}
}

func (a *App) Run() {
	commands.Execute()
}
