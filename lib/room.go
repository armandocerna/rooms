package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Room struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Cords       Cords
}

type Cords struct {
	X string `json:"x"`
	Y string `json:"y"`
}

func MakeRoom() {
	fmt.Println("vim-go")
}

func ParseRoomList() {
	rf, err := ioutil.ReadFile("lib/rooms.json")
	Check(err)

	var rooms []Room
	json.Unmarshal(rf, &rooms)
	for r := range rooms {
		fmt.Println("this is a roomlist")
		fmt.Printf("Title: %v, Description: %v, Coords: X - %d Y - %d", rooms[r].Title, rooms[r].Description, rooms[r].Cords.X, rooms[r].Cords.Y)
	}
}
