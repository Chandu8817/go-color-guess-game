package storage

import (
	"bufio"
	"encoding/json"
	"net"
	"os"
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

var ServerAddr = "localhost:8080"
var Method string = "df"
var Start = true
var Name string
var Password string
var Email string
var Age int
var Scanner = bufio.NewScanner(os.Stdin)

// Serialize the UserDetail struct into a byte slice
func SerializeUserDetail(response Response) ([]byte, error) {
	return json.Marshal(response) // Using JSON encoding

}

// Deserialize the byte slice into a UserDetail struct
func DeserializeUserDetail(data []byte) (UserDetail, error) {
	var user UserDetail
	err := json.Unmarshal(data, &user) // Using JSON decoding
	return user, err
}
