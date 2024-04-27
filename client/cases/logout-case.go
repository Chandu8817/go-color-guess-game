package cases

import (
	"fmt"
	"net"

	"github.com/color-predection/client/storage"
)

func LogoutCase(conn net.Conn) {
	inputs := storage.Scanner.Text()
	fmt.Sscan(inputs, &storage.Email)

	user := storage.UserDetail{

		Email: storage.Email,
	}

	response := storage.Response{User: user, Method: storage.Method}
	// Serialize the UserDetail struct
	serializedUser, err := storage.SerializeUserDetail(response)
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
		fmt.Println("Invalid storage.Email.")
	}

}
