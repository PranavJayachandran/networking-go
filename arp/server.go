package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)

func msgHandler(w http.ResponseWriter, r *http.Request){
	body, err := io.ReadAll(r.Body)
	fmt.Print(string(body))
	if err != nil {
		fmt.Println("Error reading the message")
	}
	io.WriteString(w, "Message Recived")
}
func sendMsgHandler(w http.ResponseWriter, r *http.Request){
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading the message")
	}

	var msg = &Message{}
	err = json.Unmarshal(body,msg)
	if !SendMessage(*msg){
		w.WriteHeader(http.StatusBadGateway)
		return 
	}
	io.WriteString(w, "Sent Successfully")
}

func arpHandler(w http.ResponseWriter, r *http.Request){
	body, err := io.ReadAll(r.Body)
	if IsError(err){
		return 
	}
	var arpMessage = &ArpMessage{}
	err = json.Unmarshal(body,arpMessage)
	if IsError(err){
		return
	}

	if arpMessage.RequestType == ArpRequest && arpMessage.RecieverName == Name{
		reply := ReplyArpRequest(arpMessage)
		w.Header().Set("Content-Type","application/json")
		json.NewEncoder(w).Encode(reply)
		return
	}
}

func ServerInit(){
	for {
		Port = ":" + strconv.Itoa(rand.Intn(5) + 8000)
		fmt.Printf("Trying to start a server at %s", Port)
		http.HandleFunc("/sendmsg", sendMsgHandler)
		http.HandleFunc("/msg", msgHandler)
		http.HandleFunc("/arp", arpHandler) 
		err := http.ListenAndServe(Port, nil)
		if err != nil {
			fmt.Print(err)
		}
	}
}
