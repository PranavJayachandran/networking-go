package main

import (
	"math/rand"
	"sync"
)


var Name string;
var Port string;

func getRandomName() string{
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 10)
	for i := range b{
		b[i] = charset[rand.Intn(len(charset))]
	}
	return "one"
	return string(b)
}

func main() {
	Name = getRandomName()
	var wg sync.WaitGroup

	wg.Add(1)

	go func(){
		defer wg.Done()
		ServerInit()
	}()

	wg.Wait()
}
