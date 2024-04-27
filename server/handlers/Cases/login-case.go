package cases

import (
	"fmt"
	"net"
	"os"

	"github.com/color-predection/server/auth"
	"github.com/color-predection/server/handlers/utils"
	"github.com/color-predection/server/storage"
)

func LoginCase(receivedData storage.Response, conn net.Conn) {
	fmt.Println("Login request")
	user, err := auth.UserLogin(receivedData.User.Email, receivedData.User.Password)

	if err != nil {
		fmt.Println("Login failed", err)

	}
	fmt.Println("login success", user)
	serializedUser, err := utils.SerializeUserDetail(user)
	if err != nil {
		fmt.Println("serializedUser failed", err)

	}

	_, err = conn.Write([]byte(serializedUser))
	if err != nil {
		fmt.Println("Error sending message:", err)
		os.Exit(1)
	}
}
