package main

import (
	"fmt"
	r "chat/api/routes"
)

func main() {
	fmt.Println("Welcome to Go Chat CLI👋")
	// Goes to routes package and brings all the required API
	r.Route()
}
