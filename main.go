package main

import (
	"fmt"
	"net/http"

	socket_io "github.com/googollee/go-socket.io"
	"github.com/leapfrog-training/go-chat/cmd"
)

func main() {
	fmt.Println("Welcome to Go Chat CLI👋")
	cmd.Execute()
	server()
}
func server() {
	server := socket_io.NewServer(nil)

	server.OnConnect("", func(s socket_io.Conn) error {
		fmt.Println("Connected: ", s.ID())
		s.Rooms()
		s.Join("chat")
		return nil
	})

	server.OnEvent("", "chat", func(s socket_io.Conn, msg string) {
		server.BroadcastToRoom("", "chat", "message", msg)
	})

	server.OnError("", func(s socket_io.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)

	fmt.Println("Serving at localhost:4000...")
	http.ListenAndServe(":4000", nil)
}
