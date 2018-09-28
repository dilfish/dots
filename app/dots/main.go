package main

import (
	"github.com/dilfish/dots"
)

func main() {
	cert := "/etc/letsencrypt/live/libsm.com-0001/fullchain.pem"
	key := "/etc/letsencrypt/live/libsm.com-0001/privkey.pem"
	ls, err := dots.GetListener(cert, key, ":853")
	if err != nil {
		panic(err)
	}
    cExit := make(chan bool)
    defer close(cExit)
	dots.Run(ls, cExit)
}
