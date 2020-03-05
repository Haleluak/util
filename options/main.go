package main

import (
	"fmt"
	"util/options/server"
	_ "util/reflect"
)

func main() {

	service := server.NewServer(
		server.Name("test Name"),
		server.Address("test Address"),
	)

	if err := service.Start(); err != nil {
		fmt.Println("loi")
	}
}

