package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Room struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Cords       struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"cords"`
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
		fmt.Printf("Title: %v, Description: %v, Coords: X - %d Y - %d\n", rooms[r].Title, rooms[r].Description, rooms[r].Cords.X, rooms[r].Cords.Y)
	}
}
