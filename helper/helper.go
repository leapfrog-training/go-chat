package helper

import (
	"fmt"
	"log"
)

/**
 * To show error if anything goes wrong or shows any error.
 * @function checkErr
 * @param {error} e
 */
func CheckErr(e error) {
	if e != nil {
		log.Println(e)
	}
}

/**
 * To print anything in the console.
 * @function printMessage
 * @param {string} message
 */
func PrintMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}