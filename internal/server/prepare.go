package server

import (
	_ "embed"
	"encoding/json"
	"github.com/candbright/wechat-robot/internal/server/db"
	"github.com/candbright/wechat-robot/internal/server/repo"
)

//go:embed res/idiom.json
var idiomByte []byte

func Prepare() error {
	var idioms []repo.Idiom
	err := json.Unmarshal(idiomByte, &idioms)
	if err != nil {
		return err
	}
	ins := db.NewDB()
	for _, idiom := range idioms {
		err = ins.AddIdiom(idiom)
		if err != nil {
			return err
		}
	}
	return nil
}
