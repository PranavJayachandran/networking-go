package main

import "fmt"

func IsError(err error) bool{
	if err != nil{
		fmt.Println(err)
		return true
	}
	return false
}
