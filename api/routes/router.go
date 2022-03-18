package routes

import (
	"log"
	"net/http"
	auth "chat/api/auth"
	// chat "chat/api/chat"
)

/**
 * All API points for the application is present here.
 * @function Route
 */
func Route() {
	// To manage routing in go.
	mux := http.NewServeMux()
	mux.HandleFunc("/login", auth.Login)
	mux.HandleFunc("/register", auth.Register)
	mux.HandleFunc("/logout", auth.Logout)
	// mux.HandleFunc("/todo", todo.Index)
	// mux.HandleFunc("/todo/create", todo.Store)
	// mux.HandleFunc("/todo/edit", todo.Update)
	// mux.HandleFunc("/todo/delete", todo.Destory)
	// mux.HandleFunc("/todo/mark-all", todo.MarkAll)

	// To start server at the host = localhost and port = 8080 with the given API endpoints.
	log.Println(http.ListenAndServe(":8080", mux))
}