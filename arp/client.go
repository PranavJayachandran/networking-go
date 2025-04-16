package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Should broadcast a address call to all the servers in the network, in this case all the ports between 8000 to 8005
func findAddress(name string)(string, bool){
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	msg := ArpMessage{
		RequestType: ArpRequest,
		SenderAddr: fmt.Sprintf("http://localhost%s", Port),
		SenderName: Name,
		RecieverAddr: "",
		RecieverName: name,
	}

	resultCh := make(chan string, 1)

	for port := 8001; port <= 8001; port++{
		go func(p int){
			url := fmt.Sprintf("http://localhost:%d/arp", p)
			jsonData, err := json.Marshal(msg)
			if IsError(err){
				return
			}

			req, err := http.NewRequestWithContext(ctx,"POST", url,bytes.NewBuffer(jsonData))

			if IsError(err){
				return 
			}

			resp, err := http.DefaultClient.Do(req)
			if IsError(err){
				return
			}

			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK{
				body, _ := io.ReadAll(resp.Body)
				fmt.Print(string(body))
				var message = &ArpMessage{}
				_ = json.Unmarshal(body, message)
				select {
				case resultCh <- string(message.SenderAddr):
					cancel()
				default:
				}
			}
		}(port)
	}

	select {
	case addr := <-resultCh:
		return addr, true
	case <-time.After(2*time.Second):
		return "", false
	}
}

func SendMessage(message Message) bool{
	addr, hasFound := findAddress(message.Name)
	if !hasFound {
		return false 
	}
	client := http.Client{Timeout: 1 * time.Second}
	data := message.Message
	req, err := http.NewRequest("POST", addr + "/msg", strings.NewReader(data))
	_, err = client.Do(req)
	if IsError(err){
		return false
	}
	return true
}

func ReplyArpRequest(request *ArpMessage) ArpMessage{
	var reply ArpMessage
	reply.SenderAddr = fmt.Sprintf("http://localhost%s", Port)
	reply.RecieverAddr = request.SenderAddr
	reply.SenderName = Name
	reply.RecieverName = request.SenderName
	reply.RequestType = ArpResponse

	return reply
}
