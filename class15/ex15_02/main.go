package main

import (
	"fmt"
	"os"

	//"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {

	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("abc: ", viper.GetString("abc"))
	fmt.Println("aaa: ", viper.GetBool("aaa"))
	fmt.Println("def: ", viper.GetString("cccccc"))
}
