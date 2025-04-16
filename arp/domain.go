package main

import "time"

type RequestType int 

const (
	ArpRequest RequestType = iota
	ArpResponse 
)

type Message struct{
	Name string 
	Message string
}

type ArpMessage struct {
	RequestType RequestType
	SenderAddr string
	RecieverAddr string
	SenderName string
	RecieverName string
}

type Cache struct {
	Value string
	Timeout time.Time
}
