package utils

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/color-predection/server/storage"
)

// Serialize the UserDetail struct into a byte slice

func SerializeUserDetail(user storage.UserDetail) ([]byte, error) {
	return json.Marshal(user) // Using JSON encoding
}

func SerializeUsers(users []storage.UserDetail) ([]byte, error) {
	return json.Marshal(users) // Using JSON encoding
}

// Deserialize the byte slice into a UserDetail struct
func DeserializeUserDetail(data []byte) (storage.Response, error) {
	var response storage.Response
	err := json.Unmarshal(data, &response) // Using JSON decoding
	return response, err
}

func SendToAllClients(message string, conn net.Conn) {

	// Send response to all clients
	storage.ClientsMux.Lock()
	for client := range storage.Clients {
		fmt.Println(client.RemoteAddr(), "rrrrrrrrr")
		if client != conn { // Skip sending response back to the sender
			_, err := client.Write([]byte(message + "!\n"))
			if err != nil {
				fmt.Println("Error sending message to client:", err)
			}
		}
	}
	storage.ClientsMux.Unlock()
}
