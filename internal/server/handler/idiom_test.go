package handler

import (
	"encoding/json"
	"fmt"
	"github.com/candbright/wechat-robot/internal/server/repo"
	"github.com/go-resty/resty/v2"
	"testing"
)

type Result[T any] struct {
	Code    int64  `json:"code"`
	Data    T      `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func TestIdiomHandler_RandomIdiom(t *testing.T) {
	endpoint := "http://localhost:11237"
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get(endpoint + "/idiom")
	if err != nil {
		t.Fatal(err)
	}
	var result Result[repo.Idiom]
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(resp.Body()))
}
