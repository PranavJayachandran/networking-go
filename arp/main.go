package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


var Name string;
var Port string;
var NameAddressCache map[string]Cache
var Timeout time.Duration

func getRandomName() string{
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 10)
	for i := range b{
		b[i] = charset[rand.Intn(len(charset))]
	}
	fmt.Print(string(b))
	return string(b)
}

func main() {
	Name = getRandomName()
	NameAddressCache = make(map[string]Cache)
	Timeout = 10 * time.Second
	var wg sync.WaitGroup

	wg.Add(1)

	go func(){
		defer wg.Done()
		ServerInit()
	}()

	wg.Wait()
}
