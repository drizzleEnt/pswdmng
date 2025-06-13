package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

/*
pass init        # Инициализация хранилища (создание файла, установка мастер-пароля)
pass login       # Вход (если используется сессия или кэширование ключа)
pass add         # Добавление новой записи
pass get         # Получение записи
pass list        # Показ всех записей
pass remove      # Удаление записи
*/

var rootcmd = &cobra.Command{
	Use:   "pass",
	Short: "password manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd")
	},
}

var initcmd = &cobra.Command{
	Use:   "init",
	Short: "initialize password manager",
	Run: func(cmd *cobra.Command, args []string) {
		psw, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return
		}
		fmt.Printf("psw: %v\n", string(psw))
		fmt.Println("init cmd")
	},
}

var logincmd = &cobra.Command{
	Use:   "login",
	Short: "login in password manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login cmd")
	},
}

var addcmd = &cobra.Command{
	Use:   "add",
	Short: "add row in password manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add cmd")
	},
}

var getcmd = &cobra.Command{
	Use:   "get",
	Short: "get row from password manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get cmd")
	},
}

var listcmd = &cobra.Command{
	Use:   "list",
	Short: "get rows list from password manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list cmd")
	},
}

var removecmd = &cobra.Command{
	Use:   "remove",
	Short: "remove row from password manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove cmd")
	},
}

func Execute() {
	if err := rootcmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func InitCommands() {
	rootcmd.AddCommand(initcmd)
	rootcmd.AddCommand(logincmd)
	rootcmd.AddCommand(addcmd)
	rootcmd.AddCommand(getcmd)
	rootcmd.AddCommand(listcmd)
	rootcmd.AddCommand(removecmd)
}
