package storage

import (
	"net"
	"sync"
)

type Response struct {
	User   UserDetail
	Method string
}

var (
	Clients    = make(map[net.Conn]struct{})
	ClientsMux sync.Mutex
)
