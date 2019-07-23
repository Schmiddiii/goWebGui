package goweb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Message The message Javascript and Go sends
type Message struct {
	ID     string   `json:"id"`
	Extras []string `json:"extras"`
}

var handlerFunc func(Message) Message

// SetUpCode sets up code
func SetUpCode(port string, handler func(Message) Message) {

	handlerFunc = handler

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/backend/respond", respondHandler)
	fmt.Println("Serving on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func respondHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	m, err := messageToObject(reqBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	response := handlerFunc(m)

	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		w.Write(nil)
		return
	}

	w.Write([]byte(responseJSON))

}

func messageToObject(msg []byte) (Message, error) {
	var m Message

	err := json.Unmarshal(msg, &m)
	if err != nil {
		return Message{}, err
	}
	return m, nil
}
