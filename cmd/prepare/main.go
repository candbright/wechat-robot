package main

import (
	"fmt"
	"github.com/candbright/wechat-robot/internal/server"
)

func main() {
	err := server.Prepare()
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
}
