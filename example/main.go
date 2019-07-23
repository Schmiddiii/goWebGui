package main

import (
	"fmt"

	goweb "github.com/Schmiddiii/goWebGui"
)

func main() {
	goweb.SetUpCode("3000", handler)
}

func handler(m goweb.Message) goweb.Message {
	fmt.Printf("Handling to Message: %+v\n", m)

	if m.ID == "btn_concat" {
		return goweb.Message{ID: "concat", Extras: []string{m.Extras[0] + m.Extras[1]}}

	}

	return goweb.Message{ID: "", Extras: []string{}}
}
