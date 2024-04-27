package cases

import (
	"fmt"
	"net"

	"github.com/color-predection/server/auth"
	"github.com/color-predection/server/game"
	"github.com/color-predection/server/handlers/utils"
	"github.com/color-predection/server/storage"
)

func CreateGameCase(receivedData storage.Response, conn net.Conn) {

	admin := auth.GetUser(receivedData.User.Email)
	if admin.IsLogin {

		if receivedData.User.Email != admin.Email {
			fmt.Println("Only admin ")
			return
		}

		create := game.CreateGame()
		if create.GameId+1 == game.Count {

			// send message to all clients

			utils.SendToAllClients("game created", conn)

		}

	}
}
