//auto generated by go-cli

package config

import (
	_ "embed"
	"github.com/candbright/go-core/config"
)

//go:embed config.yaml
var data []byte

// Config is global config parsed from config.yaml file
var Config *config.Config

func init() {
	cfg, err := config.Parse(data, config.YAML)
	if err != nil {
		panic(err)
	}
	Config = cfg
}
