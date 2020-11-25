package configs

import (
	"log"

	"github.com/BurntSushi/toml"
)

func init() {
	_, err := toml.DecodeFile("./config/config.toml", &Config)
	if err != nil {
		log.Println(err.Error())
	}
}
