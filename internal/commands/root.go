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

	r.rootCmd = factory.MakeRootCommand(func(args []string) error{
		fmt.Println(args)
		return nil
	})

	cmds := []*cobra.Command{
		factory.MakeInitCommand(r.init),
		factory.MakeAddCommand(r.add),
		factory.MakeGetCommand(r.get),
		factory.MakeListCommand(r.list),
		factory.MakeLoginCommand(r.login),
		factory.MakeRemoveCommand(r.remove),
	}
	r.cmds = cmds

	r.rootCmd.AddCommand(r.cmds...)
}
