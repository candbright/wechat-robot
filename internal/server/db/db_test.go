package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/candbright/wechat-robot/internal/server/db/options"
	"testing"
	"time"
)

func TestGetIdioms(t *testing.T) {
	ins := NewDB()
	before := time.Now()
	idioms, err := ins.GetIdioms(options.Where("word", "一劳永逸"))
	if err != nil {
		panic(err)
	}
	after := time.Now()
	marshal, err := json.Marshal(idioms)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, marshal, "", "")
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
	fmt.Println("time use:", after.UnixMilli()-before.UnixMilli(), "ms")
}
