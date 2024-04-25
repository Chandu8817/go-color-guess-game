package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

type UserDetail struct {
	Name     string
	Password string
	Email    string
	Age      int
	IsLogin  bool
}

type Response struct {
	User   UserDetail
	Method string
}

func main() {
	// Server address (change this to your server's IP and port)
	serverAddr := "localhost:8080"
	var method string
	var start = true
	var name string
	var password string
	var email string
	var age int

	// Connect to the server
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("game start")
	scanner := bufio.NewScanner(os.Stdin)
	for start {

		if !scanner.Scan() {
			break
		}

		method = scanner.Text()

		switch method {
		case "s":
			if !scanner.Scan() {
				break
			}
			inputs := scanner.Text()

			_, err := fmt.Sscan(inputs, &name, &password, &email, &age)

			if err != nil {
				fmt.Println(err)

			}
			user := UserDetail{
				Name:     name,
				Password: password,
				Email:    email,
				Age:      age,
				IsLogin:  false,
			}

			response := Response{User: user, Method: method}
			// Serialize the UserDetail struct
			serializedUser, err := serializeUserDetail(response)
			if err != nil {
				fmt.Println("Error serializing user:", err)
				return
			}

			// _, err = conn.Write([]byte(method))
			// if err != nil {
			// 	fmt.Println("Error sending met:", err)
			// 	return
			// }

			// Send the serialized user data to the server

			_, err = conn.Write(serializedUser)
			if err != nil {
				fmt.Println("Error sending message:", err)
				return
			}

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
		case "l":
			if !scanner.Scan() {
				break
			}
			inputs := scanner.Text()
			fmt.Sscan(inputs, &email, &password)

			user := UserDetail{

				Password: password,
				Email:    email,
			}

			response := Response{User: user, Method: method}
			// Serialize the UserDetail struct
			serializedUser, err := serializeUserDetail(response)
			if err != nil {
				fmt.Println("Error serializing user:", err)
				return
			}

			// _, err = conn.Write([]byte(method))
			// if err != nil {
			// 	fmt.Println("Error sending met:", err)
			// 	return
			// }

			// Send the serialized user data to the server

			_, err = conn.Write(serializedUser)
			if err != nil {
				fmt.Println("Error sending message:", err)
				return
			}

			// Read the server's response
			buffer := make([]byte, 1024)
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Error reading response:", err)
				return
			}

			res := string(buffer[:n])
			fmt.Printf("Server response: %s\n", res)

			if res != "" {
				fmt.Println("User logged in:", res)
			} else {
				fmt.Println("Invalid email/password.")
			}

		// case "ls":
		// 	users := Users
		// 	fmt.Println("User List:")
		// 	for _, user := range users {
		// 		fmt.Println(user)
		// 	}
		case "lg":
			if !scanner.Scan() {
				break
			}
			inputs := scanner.Text()
			fmt.Sscan(inputs, &email)

			user := UserDetail{

				Email: email,
			}

			response := Response{User: user, Method: method}
			// Serialize the UserDetail struct
			serializedUser, err := serializeUserDetail(response)
			if err != nil {
				fmt.Println("Error serializing user:", err)
				return
			}

			_, err = conn.Write(serializedUser)
			if err != nil {
				fmt.Println("Error sending message:", err)
				return
			}

			// Read the server's response
			buffer := make([]byte, 1024)
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Error reading response:", err)
				return
			}

			res := string(buffer[:n])
			fmt.Printf("Server response: %s\n", res)

			if res != "" {
				fmt.Println("User logout :", res)
			} else {
				fmt.Println("Invalid email.")
			}

		case "ex":
			start = false

		}

	}

}

// Serialize the UserDetail struct into a byte slice
func serializeUserDetail(response Response) ([]byte, error) {
	return json.Marshal(response) // Using JSON encoding

}

// Deserialize the byte slice into a UserDetail struct
func deserializeUserDetail(data []byte) (UserDetail, error) {
	var user UserDetail
	err := json.Unmarshal(data, &user) // Using JSON decoding
	return user, err
}
