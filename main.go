package main

import (
	"fmt"
	"github.com/armandocerna/rooms/lib"
)

func main() {
	fmt.Println("This is a room")
	// fmt.Printf("%s", lib.MakeRoom("HOME"))
	lib.ParseRoomList()
}
