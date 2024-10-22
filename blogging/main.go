package main

import (
	"fmt"
	"matwa/blogger/server"
)

func main() {
	sv := server.NewServer(":8800")
	if err := sv.Start(); err != nil {
		panic(err)
	}
	fmt.Println("Server started")
}
