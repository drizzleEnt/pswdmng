package factory

import (
	"github.com/spf13/cobra"
)

func NewCommandFactory() *CommandFactory {
	return &CommandFactory{}
}

type CommandFactory struct {
}

func (f *CommandFactory) MakeRootCommand(incomeFunc func(cmd *cobra.Command, args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pass",
		Short: "password manager",
		Run:   incomeFunc,
	}

	return cmd
}

func (f *CommandFactory) MakeInitCommand(incomeFunc func(cmd *cobra.Command, args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize password store",
		Run:   incomeFunc,
	}

	return cmd
}

func (f *CommandFactory) MakeGetCommand(incomeFunc func(cmd *cobra.Command, args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get row from password manager",
		Run:   incomeFunc,
	}

	return cmd
}

func (f *CommandFactory) MakeListCommand(incomeFunc func(cmd *cobra.Command, args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "get rows list from password manager",
		Run:   incomeFunc,
	}

	return cmd
}

func (f *CommandFactory) MakeRemoveCommand(incomeFunc func(cmd *cobra.Command, args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "remove row from password manager",
		Run:   incomeFunc,
	}

	return cmd
}

func (f *CommandFactory) MakeLoginCommand(incomeFunc func(cmd *cobra.Command, args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "login in password manager",
		Run:   incomeFunc,
	}

	return cmd
}

func (f *CommandFactory) MakeAddCommand(incomeFunc func(cmd *cobra.Command, args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add row in password manager",
		Run:   incomeFunc,
	}

	return cmd
}
