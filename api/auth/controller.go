package auth

import (
	config "chat/api/config"
	helper "chat/api/helper"
	"encoding/json"
	"net/http"
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

func Login(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/login" {
		http.NotFound(w, r)
		return
	}
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")

		var res User

		err := json.NewDecoder(r.Body).Decode(&res)
		helper.CheckErr(err)

		email := res.Email
		password := res.Password
		username := res.Username

		if email == "" || password == "" || username == "" {
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(&Response{"failure", 204, "Empty field, No data sent."})
		} else {
			db := config.SetupDB()

			row := db.QueryRow("SELECT COUNT(username) As tot_username FROM users WHERE username=$1 AND pwd=$2", username,password)

			var usernameExist int
			err = row.Scan(&usernameExist)

			helper.CheckErr(err)

			if usernameExist <= 0 {
				w.WriteHeader(http.StatusNotAcceptable)
				json.NewEncoder(w).Encode(&Response{"failure", 406, "email or password not correct"})
				return
			}

			db.Close()

			IsLogin=true

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&Response{"success", 200, IsLogin})
		}
	} else {
		http.Error(w, "Only POST requests are allowed!", http.StatusMethodNotAllowed)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		http.NotFound(w, r)
		return
	}
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")

		var res User

		err := json.NewDecoder(r.Body).Decode(&res)
		helper.CheckErr(err)

		email := res.Email
		password := res.Password
		username := res.Username

		if email == "" || password == "" || username == "" {
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(&Response{"failure", 204, "Empty field, No data sent."})
		} else {
			db := config.SetupDB()

			row := db.QueryRow("SELECT COUNT(email) As tot_email FROM users WHERE email=$1 OR username=$2;", email,username)

			var emailExist int
			err = row.Scan(&emailExist)

			helper.CheckErr(err)

			if emailExist > 0 {
				w.WriteHeader(http.StatusNotAcceptable)
				json.NewEncoder(w).Encode(&Response{"failure", 406, "Email already exist"})
				return
			}

			var lastInsertId int

			err := db.QueryRow("INSERT INTO users(username,email,pwd) VALUES($1, $2, $3) returning id;", username, email, password).Scan(&lastInsertId)

			helper.CheckErr(err)
			db.Close()

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&Response{"success", 200, lastInsertId})
		}

	} else {

		http.Error(w, "Only POST requests are allowed!", http.StatusMethodNotAllowed)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/logout" {
		http.NotFound(w, r)
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
			IsLogin=false
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&Response{"success", 200, IsLogin})
	} else {
		http.Error(w, "Only POST requests are allowed!", http.StatusMethodNotAllowed)
	}
}