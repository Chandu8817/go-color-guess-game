package game

import (
	"fmt"
	"time"

	"github.com/color-predection/server/storage"
)

var Count int = 1

func CreateGame() storage.GameDetial {

	fmt.Println("Creating new a game")

	start := time.Now().Add(time.Second * 60)

	game := storage.GameDetial{GameId: Count, Colors: [2]string{"Blue", "Red"}, BidStartTime: start, BidtEndTime: start.Add(time.Minute * 5)}
	fmt.Println(game)

	storage.GameList = append(storage.GameList, game)
	Count++

	return game

}
