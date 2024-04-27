package main

import (
	"fmt"
	"net"
	"os"

	"github.com/color-predection/server/handlers"
	"github.com/color-predection/server/storage"
)

func main() {

	args := os.Args[1:]

	if len(args) > 0 && args[0] == "createAdmin" {
		// Print the arguments
		fmt.Println("Arguments provided:", args)

		// Access individual arguments

		admin := storage.UserDetail{Name: "admin", Password: "admin", Email: "admin@gmail.com", Age: 25, IsLogin: false}
		storage.Users = append(storage.Users, admin)

	} else {
		fmt.Println("No arguments provided.")
	}

	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Print("Error listening : ", err.Error())
		return

	}
	defer ln.Close()
	fmt.Printf("TCP server listening on port %d...\n", 8080)

	for {

		conn, err := ln.Accept()
		storage.ClientsMux.Lock()

		fmt.Println("New client connected:", conn.RemoteAddr())
		storage.ClientsMux.Unlock()

		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handlers.HandleConnection(conn)

	}
}
