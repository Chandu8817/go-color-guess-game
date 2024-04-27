package cases

import (
	"fmt"
	"net"
	"os"

	"github.com/color-predection/server/auth"
	"github.com/color-predection/server/handlers/utils"
	"github.com/color-predection/server/storage"
)

func SignUpCase(receivedData storage.Response, conn net.Conn) {
	fmt.Println("Received users:")
	auth.UserSignUp(receivedData.User)
	for i := 0; i < len(storage.Users); i++ {

		fmt.Println(storage.Users[i])
	}

	// Serialize the UserDetail struct
	serializedUser, err := utils.SerializeUserDetail(receivedData.User)
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

func UserListCase(receivedData storage.Response, conn net.Conn) {

	users := auth.UserList()
	_serializeUsers, _ := utils.SerializeUsers(users)
	_, err := conn.Write([]byte(_serializeUsers))

	if err != nil {
		fmt.Println("Error sending message:", err)
		os.Exit(1)
	}

}
