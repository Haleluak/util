package options

import (
	"fmt"
	"tutorial/options/server"
	_ "tutorial/reflect"
)

func test() {

	service := server.NewServer(
		server.Name("test Name"),
		server.Address("test Address"),
	)

	if err := service.Start(); err != nil {
		fmt.Println("loi")
	}
}

