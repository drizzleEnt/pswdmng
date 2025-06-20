package factory

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func MakeRootCommand(incomeFunc func(args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pass",
		Short: "password manager",
		Run: func(cmd *cobra.Command, args []string) {
			incomeFunc(args)
		},
	}

	return cmd
}

func MakeInitCommand(incomeFunc func(bool)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize password store",
		Run: func(cmd *cobra.Command, args []string) {
			newFlag, err := cmd.Flags().GetBool("new")
			if err != nil {
				fmt.Printf("Error: %v\n", err.Error())
				os.Exit(1)
			}
			incomeFunc(newFlag)
		},
	}
	usage := "Create new passwords account"
	cmd.Flags().BoolP("new", "n", false, usage)

	return cmd
}

func MakeGetCommand(incomeFunc func([]string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get row from password manager",
		Run: func(cmd *cobra.Command, args []string) {
			incomeFunc(args)
		},
	}

	return cmd
}

func MakeListCommand(incomeFunc func(args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "get rows list from password manager",
		Run: func(cmd *cobra.Command, args []string) {
			incomeFunc(args)
		},
	}

	return cmd
}

func MakeRemoveCommand(incomeFunc func(args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "remove row from password manager",
		Run: func(cmd *cobra.Command, args []string) {
			incomeFunc(args)
		},
	}

	return cmd
}

func MakeLoginCommand(incomeFunc func(args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "login in password manager",
		Run: func(cmd *cobra.Command, args []string) {
			incomeFunc(args)
		},
	}

	return cmd
}

func MakeAddCommand(incomeFunc func(string, string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add row in password manager",
		Run: func(cmd *cobra.Command, args []string) {
			login, err := cmd.Flags().GetString("login")
			if err != nil {
				fmt.Printf("err.Error(): %v\n", err.Error())
				os.Exit(1)
			}

			url, err := cmd.Flags().GetString("url")
			if err != nil {
				fmt.Printf("err.Error(): %v\n", err.Error())
				os.Exit(1)
			}

			incomeFunc(login, url)
		},
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
