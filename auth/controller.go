package auth

import (
	"fmt"
	"log"

	config "github.com/leapfrog-training/go-chat/config"
	helper "github.com/leapfrog-training/go-chat/helper"
)

/**
 * To structure the response given by the application
 * @type {struct} Response
 */
type Response struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data,omitempty"`
}

var IsLogin bool

func Login(username,email, password string) (User, error)  {
			db := config.SetupDB()

			row := db.QueryRow("SELECT * FROM users WHERE username=$1 AND email=$2 AND pwd=$3", username,email,password)

			var id int
			var dbUsername, dbEmail, dbPassword string
			err := row.Scan(&id,&dbUsername,&dbEmail,&dbPassword)

			helper.CheckErr(err)

			db.Close()

			return User{Username: username, Email: email, Password: password},err
}

func Register(username,email, password string) {
			db := config.SetupDB()

			row := db.QueryRow("SELECT COUNT(email) As tot_email FROM users WHERE email=$1 OR username=$2;", email,username)

			var emailExist int
			err := row.Scan(&emailExist)

			helper.CheckErr(err)

			if emailExist > 0 {
				log.Fatalf("Connection error")
			}

			var lastInsertId int

			err = db.QueryRow("INSERT INTO users(username,email,pwd) VALUES($1, $2, $3) returning id;", username, email, password).Scan(&lastInsertId)

			helper.CheckErr(err)
			db.Close()
			fmt.Println("Register successful!")
}

func ChatStore(username,chat string) {
	db := config.SetupDB()

	var lastInsertId int

	err := db.QueryRow("INSERT INTO chat(username,chat) VALUES($1, $2) returning id;", username,chat).Scan(&lastInsertId)

	helper.CheckErr(err)
	db.Close()
	fmt.Println("Chat stored successful!")
}
