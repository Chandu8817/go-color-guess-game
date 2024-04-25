package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/color-predection/server/storage"
)

type Response struct {
	User   storage.UserDetail
	Method string
}

func main() {

	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Print("Error listening : ", err.Error())
		return

	}
	defer ln.Close()
	fmt.Printf("TCP server listening on port %d...\n", 8080)
	for {

		conn, err := ln.Accept()

		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		// Deserialize the byte slice into a UserDetail struct and method

		receivedData, err := deserializeUserDetail(buffer[:n])

		if err != nil {
			fmt.Println("Error deserializing user:", err)
			return
		}

		// Send a response to the client and print the response on server
		if receivedData.Method == "s" {
			storage.Users = append(storage.Users, receivedData.User)
			// Use receivedUser, which is of type UserDetail
			fmt.Println("Received users:")
			for i := 0; i < len(storage.Users); i++ {

				fmt.Println(storage.Users[i])
			}

			// Serialize the UserDetail struct
			serializedUser, err := serializeUserDetail(receivedData.User)
			if err != nil {
				fmt.Println("Error serializing user:", err)
				return
			}
			_, err = conn.Write([]byte(serializedUser))
			if err != nil {
				fmt.Println("Error sending message:", err)
				os.Exit(1)
			}

		}

		if receivedData.Method == "l" {

		}

	}

}

// Serialize the UserDetail struct into a byte slice
func serializeUserDetail(user storage.UserDetail) ([]byte, error) {
	return json.Marshal(user) // Using JSON encoding
}

// Deserialize the byte slice into a UserDetail struct
func deserializeUserDetail(data []byte) (Response, error) {
	var response Response
	err := json.Unmarshal(data, &response) // Using JSON decoding
	return response, err
}
