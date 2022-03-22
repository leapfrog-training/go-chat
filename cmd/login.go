package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"


	"github.com/leapfrog-training/go-chat/auth"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

var User auth.User

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login user",
	Long:  `Login user with username, email and password`,
	Run: func(cmd *cobra.Command, args []string) {
		systemLogin()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func promptGetInput(pc string) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

func systemLogin() {
	username := promptGetInput("Enter Username: ")

	email := promptGetInput("Enter Email: ")

	password := promptGetInput("Enter Password: ")

	user, err := auth.Login(username, email, password)

	if err != nil {
		fmt.Println("Email or Password is incorrect. Try again")
		systemLogin()
	}
	User = user
	fmt.Println("Login: ", User)
	
	client()

}


func client() {
	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}

	opts.Query["user"] = User.Email
	opts.Query["pwd"] = User.Password
	uri := "http://localhost:4000/"

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		fmt.Printf("NewClient error:%v\n", err)
		return
	}

	client.On("error", func() {
		fmt.Printf("on error\n")
	})

	client.On("connection", func() {
		fmt.Println("Client ID: ")
	})
	client.On("message", func(msg string) {
		fmt.Println("Received: ", msg)
	})
	client.On("disconnection", func() {
		fmt.Printf("on disconnect\n")
	})

	fmt.Println("Chat room started: ")
	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		command := string(data)
		client.Emit("chat", command)
		auth.ChatStore(User.Username, command)
		fmt.Printf("Send: %v\n", command)
	}
}


