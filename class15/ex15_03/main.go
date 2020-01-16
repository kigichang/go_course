package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// LoadFile ...
func LoadFile(config string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigFile(config)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}

func main() {

	config, err := LoadFile("myconfig.json")
	if err != nil {
		log.Println(err)
		return
	}

	if err := config.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("abc: ", config.GetString("abc"))
	fmt.Println("aaa: ", config.GetBool("aaa"))
	fmt.Println("def: ", config.GetString("cccccc"))
}
