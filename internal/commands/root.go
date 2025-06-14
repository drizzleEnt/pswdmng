package commands

import (
	"fmt"
	"os"
	"pswdmng/internal/commands/factory"
	"pswdmng/internal/repository"

	"github.com/spf13/cobra"
)

func New(r repository.Repository) *Root {
	root := &Root{
		repo: r,
	}

	return root
}

type Root struct {
	repo    repository.Repository
	rootCmd *cobra.Command
	cmds    []*cobra.Command
}

func (r *Root) Execute() {
	if err := r.rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func (r *Root) InitCommands() {
	f := factory.NewCommandFactory()

	r.rootCmd = f.MakeRootCommand(func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd")
	})

	cmds := []*cobra.Command{
		f.MakeInitCommand(r.init),
		f.MakeAddCommand(r.add),
		f.MakeGetCommand(r.get),
		f.MakeListCommand(r.list),
		f.MakeLoginCommand(r.login),
		f.MakeRemoveCommand(r.remove),
	}
	r.cmds = cmds

	r.rootCmd.AddCommand(r.cmds...)
}
