package commands

import (
	"fmt"
	"os"
	"pswdmng/internal/repository/dbrepo"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var initcmd = &cobra.Command{
	Use:   "init",
	Short: "initialize password manager",
	Run: func(cmd *cobra.Command, args []string) {
		psw, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return
		}
		//check exit 
		r := dbrepo.New()
		r.CheckExist()
		fmt.Println("")
		fmt.Printf("psw: %v\n", string(psw))
		fmt.Println("init cmd")
	},
}
