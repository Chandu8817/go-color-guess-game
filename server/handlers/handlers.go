package handlers

import (
	"fmt"
	"net"

	cases "github.com/color-predection/server/handlers/Cases"
	"github.com/color-predection/server/handlers/utils"
	"github.com/color-predection/server/storage"
)

func HandleConnection(conn net.Conn) {

	defer func() {
		// Remove client from the list when it disconnects
		storage.ClientsMux.Lock()
		delete(storage.Clients, conn)
		storage.ClientsMux.Unlock()
		conn.Close()
	}()

	storage.ClientsMux.Lock()
	storage.Clients[conn] = struct{}{}
	storage.ClientsMux.Unlock()

	// handleClient(conn)

	for {

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		// Deserialize the byte slice into a UserDetail struct and method

		receivedData, err := utils.DeserializeUserDetail(buffer[:n])

		if err != nil {
			fmt.Println("Error deserializing user:", err)
			return
		}

		// Send a response to the client and print the response on server
		if receivedData.Method == "s" {

			cases.SignUpCase(receivedData, conn)

		}
		if receivedData.Method == "l" {

			cases.LoginCase(receivedData, conn)

		}
		if receivedData.Method == "lg" {
			fmt.Println("receivedData", receivedData.Method)
			cases.LogoutCase(receivedData, conn)

		}
		if receivedData.Method == "ls" {
			cases.UserListCase(receivedData, conn)

		}
		if receivedData.Method == "cr" {

			cases.CreateGameCase(receivedData, conn)

		}

	}

}
