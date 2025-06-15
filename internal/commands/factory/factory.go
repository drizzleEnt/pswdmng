package factory

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func MakeRootCommand(incomeFunc func(cmd *cobra.Command, args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pass",
		Short: "password manager",
		Run:   incomeFunc,
	}

	return cmd
}

func MakeInitCommand(incomeFunc func(cmd *cobra.Command, args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize password store",
		Run:   incomeFunc,
	}
	usage := "Create new passwords account"
	cmd.Flags().BoolP("new", "n", false, usage)

	return cmd
}

func MakeGetCommand(incomeFunc func(cmd *cobra.Command, args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get row from password manager",
		Run:   incomeFunc,
	}

	return cmd
}

func MakeListCommand(incomeFunc func(cmd *cobra.Command, args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "get rows list from password manager",
		Run:   incomeFunc,
	}

	return cmd
}

func MakeRemoveCommand(incomeFunc func(cmd *cobra.Command, args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "remove row from password manager",
		Run:   incomeFunc,
	}

	return cmd
}

func MakeLoginCommand(incomeFunc func(cmd *cobra.Command, args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "login in password manager",
		Run:   incomeFunc,
	}

	return cmd
}

func MakeAddCommand(incomeFunc func(cmd *cobra.Command, args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add row in password manager",
		Run:   incomeFunc,
	}

	usage := "Your new login"
	cmd.Flags().StringP("login", "l", "", usage)
	if err := cmd.MarkFlagRequired("login"); err != nil {
		fmt.Println("failed mark login flag required")
		os.Exit(1)
	}

	usage = "URL of site to add new row"
	cmd.Flags().StringP("url", "u", "", usage)
	if err := cmd.MarkFlagRequired("url"); err != nil {
		fmt.Println("failed mark url flag required")
		os.Exit(1)
	}

	return cmd
}
