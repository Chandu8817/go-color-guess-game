package main

import (
	"fmt"
	"net"
	"os"
	"time"

	cases "github.com/color-predection/client/cases"
	"github.com/color-predection/client/storage"
)

func main() {
	// Server address (change this to your server's IP and port)

	// Connect to the server
	conn, err := net.Dial("tcp", storage.ServerAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("game start")

	for storage.Start {

		if !storage.Scanner.Scan() {
			break
		}

		storage.Method = storage.Scanner.Text()

		switch storage.Method {
		case "s":
			if !storage.Scanner.Scan() {
				break
			}
			cases.SignUpCase(conn)

		case "l":
			if !storage.Scanner.Scan() {
				break
			}
			cases.LoginCase(conn)
		case "ls":
			cases.UserListCase(conn)

		case "lg":
			if !storage.Scanner.Scan() {
				break
			}
			cases.LogoutCase(conn)
		case "ex":
			storage.Start = false

		case "cr":
			if !storage.Scanner.Scan() {
				break
			}
			cases.CreateGameCase(conn)

		default:
			// Read the server's response
			buffer := make([]byte, 1024)
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Error reading response:", err)
				return
			}

			res := string(buffer[:n])
			fmt.Printf("Server response: %s\n", res)

			// Sleep for some time before sending the next message
			time.Sleep(5 * time.Second)

		}

	}

}
