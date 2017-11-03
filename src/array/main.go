package main

import (
	"encoding/json"
	"fmt"
)

type GameMapType [4][4]uint32

func (m *GameMapType) ToString() (strMap string, err error) {
	byteJson, err := json.Marshal(m)
	if err != nil {
		return
	}

	strMap = string(byteJson)
	return

}

func (m *GameMapType) ParseString(strMap string) (err error) {
	err = json.Unmarshal([]byte(strMap), m)
	return
}

func NewGameMap(strMap string) (gameMap *GameMapType, err error) {
	gameMap = &GameMapType{}
	err = gameMap.ParseString(strMap)
	return
}

func main() {
	var gameMap = GameMapType{
		{0,0,0,0},
		{0,0,0,0},
		{0,0,0,1},
		{0,0,0,1},
	}

	strGameMap, _ := gameMap.ToString()
	fmt.Println(strGameMap)

	fmt.Println(NewGameMap(strGameMap))

}
