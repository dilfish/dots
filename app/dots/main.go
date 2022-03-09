// Copyright 2018 Sean.ZH

package main

import (
	"github.com/dilfish/dots"
	"log"
)

func main() {
	cert := "/etc/letsencrypt/live/dilfish.dev-0001/fullchain.pem"
	key := "/etc/letsencrypt/live/dilfish.dev-0001/privkey.pem"
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ls, err := dots.GetListener(cert, key, ":853")
	if err != nil {
		panic(err)
	}
	cExit := make(chan bool)
	defer close(cExit)
	dots.Run(ls, cExit)
}
