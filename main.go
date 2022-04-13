package main

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("app/config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

}

func main() {

}
