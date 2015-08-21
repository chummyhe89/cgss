// player
package cg

import (
	"fmt"
)

type Player struct {
	Name  string
	Level string
	Exp   string
	Room  int

	mq chan *Message //waiting to receive the massage
}

func NewPlayer() *Player {
	m := make(chan *Message, 1024)
	player := &Player("", 0, 0, 0, m)
	go func(p *Player) {
		for {
			msg := <-p.mq
			fmt.Println(p.Name, "receviced message:", msg.Content)
		}
	}(player)
	return player
}

//func main() {
//	fmt.Println("Hello World!")
//}
