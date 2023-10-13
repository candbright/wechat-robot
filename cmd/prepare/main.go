package main

import (
	"fmt"
	"github.com/candbright/wechat-robot/internal/prepare"
)

func main() {
	err := prepare.IdiomPrepare()
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
}
