package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	goweb "github.com/Schmiddiii/goWebGui"
)

var field [3][3]string
var player = 0
var finished = false

func main() {
	goweb.SetUpCode("3001", handler)
}

func handler(msg goweb.Message) goweb.Message {

	// Resetting
	if msg.ID == "reset" {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				field[i][j] = ""
			}
		}
		player = 0
		finished = false
		return goweb.Message{ID: "clear", Extras: []string{}}
	}

	// Do nothing if already won
	if finished {
		return goweb.Message{}
	}
	// Calculate move
	fmt.Printf("Message: %+v\n", msg)

	y, err := strconv.Atoi(string(msg.ID[0]))
	x, err := strconv.Atoi(string(msg.ID[2]))

	if err != nil {
		fmt.Println(err)
		return goweb.Message{}
	}
	switch field[x][y] {
	case "":
		field[x][y] = strconv.Itoa(player)
		player = (player + 1) % 2
		break
	default:
		break
	}

	fiByte, err := json.Marshal(field)
	if err != nil {
		fmt.Println(err)
		return goweb.Message{}
	}

	fiStr := string(fiByte)

	// WinTests
	if ((field[0][0] == field[1][1] && field[0][0] == field[2][2]) || (field[2][0] == field[1][1] && field[2][0] == field[0][2])) && field[1][1] != "" {
		finished = true
		return goweb.Message{ID: "win", Extras: []string{fiStr, strconv.Itoa((player + 1) % 2)}}
	}
	for i := 0; i < 3; i++ {
		if ((field[i][0] == field[i][1] && field[i][0] == field[i][2]) || (field[0][i] == field[1][i] && field[0][i] == field[2][i])) && field[i][i] != "" {
			finished = true
			return goweb.Message{ID: "win", Extras: []string{fiStr, strconv.Itoa((player + 1) % 2)}}
		}
	}

	tie := true
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if field[i][j] == "" {
				tie = false
			}
		}
	}

	if tie {
		return goweb.Message{ID: "tie", Extras: []string{fiStr}}
	}

	return goweb.Message{ID: "move", Extras: []string{fiStr}}
}
