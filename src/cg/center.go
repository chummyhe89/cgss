// center
package cg

import (
	"encoding/json"
	"errors"
	"sync"
	"gogame/cgss/src/ipc"
)

var _ ipc.Server = &CenterServer()  //确认实现了Server借口

type Message struct {
	From string `json:"from”`
	To string `json:"to“`
	Content string `json:"content"`
}

type CenterServer struct {
	servers map[string] ipc.Server
	players []*Player
	room []*Room
	mutex sync.RWMutex
}
func
//func main() {
//	fmt.Println("Hello World!")
//}
