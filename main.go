package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"time"
)

func MakeClient() (net.Conn, error) {
	cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
	if err != nil {
		fmt.Println("load key", err)
		return nil, err
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: false}
	return tls.Dial("tcp", "1.1.1.1:853", &config)
	// debug
	// state := conn.ConnectionState()
	// for _, v := range state.PeerCertificates
	// v.Subject
	// x509.MarshalPKIXPublicKey(v.PublicKey)
	// state.HandshakeComplete
	// state.NegotiatedProtocolIsMutual
}

func GetLisener(cert, key string) (net.Listener, error) {
	var err error
	config := &tls.Config{}
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0], err = tls.LoadX509KeyPair(cert, key)
	if err != nil {
		fmt.Println("load key pair", err)
		return nil, err
	}
	config.BuildNameToCertificate()
	conn, err := net.Listen("tcp", ":853")
	if err != nil {
		fmt.Println("listen 853", err)
		return nil, err
	}
	ls := tls.NewListener(conn, config)
	return ls, nil
}

func HandleAC(conn net.Conn) {
	defer conn.Close()
	fmt.Println("time.Now", time.Now(), conn.RemoteAddr())
	cli, err := MakeClient()
	if err != nil {
		fmt.Println("get client", err)
		return
	}
	defer cli.Close()
	// ignore error
	go io.Copy(cli, conn)
	io.Copy(conn, cli)
}

func Run(ls net.Listener) {
	defer ls.Close()
	for {
		ac, err := ls.Accept()
		if err != nil {
			fmt.Println("accept error", err)
			break
		}
		go HandleAC(ac)
	}
}

func main() {
	cert := "/etc/letsencrypt/live/libsm.com-0001/fullchain.pem"
	key := "/etc/letsencrypt/live/libsm.com-0001/privkey.pem"
	ls, err := GetLisener(cert, key)
	if err != nil {
		panic(err)
	}
	Run(ls)
}
