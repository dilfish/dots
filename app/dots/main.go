package main

import (
	"github.com/dilfish/dots"
)

func main() {
	cert := "/etc/letsencrypt/live/libsm.com-0001/fullchain.pem"
	key := "/etc/letsencrypt/live/libsm.com-0001/privkey.pem"
	ls, err := dots.GetLisener(cert, key)
	if err != nil {
		panic(err)
	}
	dots.Run(ls)
}
