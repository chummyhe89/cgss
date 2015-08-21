// center
package cg

import (
	"encoding/json"
	"errors"
	"sync"

	"gogame/cgss/src/ipc"
)

var _ ipc.Server = &CenterServer()

//func main() {
//	fmt.Println("Hello World!")
//}
