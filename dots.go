// Copyright 2018 Sean.ZH

package dots

import (
	"crypto/tls"
	"io"
	"net"
	"log"
)

// MakeClient create a tcp client
func MakeClient() (net.Conn, error) {
	pem := "/root/go/src/github.com/dilfish/dots/testdata/certs/client.pem"
	key := "/root/go/src/github.com/dilfish/dots/testdata/certs/client.key"
	cert, err := tls.LoadX509KeyPair(pem, key)
	if err != nil {
		log.Println("load key", err)
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

// GetListener create an tls client
func GetListener(cert, key, portStr string) (net.Listener, error) {
	var err error
	config := &tls.Config{}
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0], err = tls.LoadX509KeyPair(cert, key)
	if err != nil {
		log.Println("load key pair", err)
		return nil, err
	}
	config.BuildNameToCertificate()
	conn, err := net.Listen("tcp", portStr)
	if err != nil {
		log.Println("listen 853", err)
		return nil, err
	}
	ls := tls.NewListener(conn, config)
	return ls, nil
}

func handleAC(conn net.Conn) {
	defer conn.Close()
	log.Println("new conn:", conn.RemoteAddr())
	cli, err := MakeClient()
	if err != nil {
		log.Println("get client", err)
		return
	}
	defer cli.Close()
	// ignore error
	log.Println("conn is", conn)
	go io.Copy(cli, conn)
	io.Copy(conn, cli)
	bt := make([]byte, 100)
	n, err := conn.Read(bt)
	log.Println("we read", string(bt), n, err)
}

func doLs(ls net.Listener, c chan net.Conn) {
	for {
		ac, err := ls.Accept()
		if err != nil {
			ls.Close()
			c <- nil
			close(c)
			return
		}
		c <- ac
	}
}

// Run loops listen on tls
func Run(ls net.Listener, cExit chan bool) {
	cConn := make(chan net.Conn)
	go doLs(ls, cConn)
	for {
		select {
		case conn := <-cConn:
			if conn == nil {
				return
			}
			go handleAC(conn)
		case <-cExit:
			ls.Close()
		}
	}
}
