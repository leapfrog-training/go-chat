package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	helper "github.com/leapfrog-training/go-chat/helper"
)


/**
 * Creating Global var to be accessed from outside the package.
 */
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "chatdb"
)

/**
 * Setup for Database and Initializing Collection for it to be used.
 * @function setup
 */
func SetupDB() *sql.DB {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbInfo)

	helper.CheckErr(err)

	helper.PrintMessage("Connected!")
	return db
}
