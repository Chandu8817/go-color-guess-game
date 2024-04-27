package cases

import (
	"fmt"
	"net"
	"os"

	"github.com/color-predection/server/auth"
	"github.com/color-predection/server/handlers/utils"
	"github.com/color-predection/server/storage"
)

func LogoutCase(receivedData storage.Response, conn net.Conn) {
	fmt.Println("Logout request")
	logout := auth.LogOut(receivedData.User.Email)
	fmt.Println(logout, "xxxcv")

	if logout != "" {
		user := auth.GetUser(receivedData.User.Email)
		fmt.Println("Logout success", user)
		serializedUser, err := utils.SerializeUserDetail(user)
		if err != nil {
			fmt.Println("serializedUser failed", err)

		}

		_, err = conn.Write(serializedUser)
		if err != nil {
			fmt.Println("Error sending message:", err)
			os.Exit(1)
		}
	}
}
