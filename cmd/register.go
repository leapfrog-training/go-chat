package cmd

import (
	"github.com/leapfrog-training/go-chat/auth"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register Here: ",
	Long:  `Register new user`,
	Run: func(cmd *cobra.Command, args []string) {
		registerUser()
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}

func registerUser() {
	name := promptGetInput("Enter Name: ")

	email := promptGetInput("Enter Email: ")

	password := promptGetInput("Enter Password: ")

	auth.Register(name, email, password)

	systemLogin()
}