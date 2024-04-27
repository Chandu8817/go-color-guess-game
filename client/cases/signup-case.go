package cases

import (
	"fmt"
	"net"
	"time"

	"github.com/color-predection/client/storage"
)

func SignUpCase(conn net.Conn) {

	inputs := storage.Scanner.Text()

	_, err := fmt.Sscan(inputs, &storage.Name, &storage.Password, &storage.Email, &storage.Age)

	if err != nil {
		fmt.Println(err)

	}
	user := storage.UserDetail{
		Name:     storage.Name,
		Password: storage.Password,
		Email:    storage.Email,
		Age:      storage.Age,
		IsLogin:  false,
	}

	response := storage.Response{User: user, Method: storage.Method}
	// Serialize the UserDetail struct
	serializedUser, err := storage.SerializeUserDetail(response)
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
}

func UserListCase(conn net.Conn) {

	serializeUser, _ := storage.SerializeUserDetail(storage.Response{User: storage.UserDetail{Name: ""}, Method: storage.Method})
	conn.Write(serializeUser)

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("failed to read", err)

	}
	users := string(buffer[:n])

	fmt.Println("User List:", users)
	// for _, user := range users {
	// 	fmt.Println(user)
	// }

}
