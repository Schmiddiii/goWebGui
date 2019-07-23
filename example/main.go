package main

import (
	"fmt"

	goweb "github.com/Schmiddiii/goWeb"
)

func main() {
	goweb.SetUpCode("3000", handler)
}

func handler(m goweb.Message) goweb.Message {
	fmt.Printf("Handling to Message: %+v\n", m)

	if m.ID == "btn" && m.Extras[0] == "asd" {
		fmt.Println("You clicked btn 1 with first extra asd")
	}

	return goweb.Message{ID: "adslköjh", Extras: []string{"Hi", "asdo"}}
}
