package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	name  string
	proxy bool
	test  string
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rootCmd := &cobra.Command{Use: "myapp"}

	createCmd := &cobra.Command{Use: "create"}

	updateCmd := &cobra.Command{Use: "update"}

	createCmd.Flags().StringVarP(&name, "name", "n", "myname", "assign a name")
	createCmd.Flags().BoolVarP(&proxy, "proxy", "p", false, "use proxy to connect")

	createCmd.Args = cobra.ExactArgs(1)

	createCmd.Run = func(cmd *cobra.Command, args []string) {
		fmt.Println("creating")
		fmt.Println("name:", name)
		fmt.Println("proxy:", proxy)
		fmt.Println("args:", args)
	}

	updateCmd.Run = func(cmd *cobra.Command, args []string) {
		fmt.Println("viper test:", viper.GetString("test"))
		fmt.Println(args)
	}

	rootCmd.PersistentFlags().StringVarP(&test, "test", "t", "my test", "test string")
	viper.BindPFlag("test", rootCmd.PersistentFlags().Lookup("test"))

	rootCmd.AddCommand(createCmd, updateCmd)

	rootCmd.Execute()
}
