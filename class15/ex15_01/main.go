package main

import (
	"flag"
	"log"
)

var (
	account, password string
	debug             bool
)

func main() {
	flag.StringVar(&account, "account", account, "account to login")
	flag.StringVar(&password, "password", password, "password for account")
	flag.BoolVar(&debug, "debug", debug, "dump account and password or not")

	flag.Parse()

	if account == "" || password == "" {
		flag.PrintDefaults()
		return
	}

	if debug {
		log.Println("account:", account, "password:", password)
	}

	log.Println("end")
}
