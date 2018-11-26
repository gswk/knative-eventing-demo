package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	ID          string `json:"ID"`
	Data        string `json:"Data"`
	PublishTime string `json:"PublishTime"`
}

func handlePost(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)

	var message Message
	json.Unmarshal(body, &message)
	sData, _ := base64.StdEncoding.DecodeString(message.Data)

	fmt.Fprintf(rw, "[%s %s] %s", message.PublishTime, message.ID, sData)
	log.Printf("[%s %s] %s", message.PublishTime, message.ID, sData)
}

func main() {
	log.Print("Starting server on port 8080...")
	http.HandleFunc("/", handlePost)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
