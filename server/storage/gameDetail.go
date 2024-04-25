package storage

import "time"

type GameDetial struct {
	GameId       int
	Players      []UserDetail
	PlayerBid    map[UserDetail]int
	Winner       UserDetail
	Colors       [2]string
	BidStartTime time.Time
	BidtEndTime  time.Time
	Result       string
}

var GameList []GameDetial
