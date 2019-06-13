// Copyright 2018 Sean.ZH

package dots

import (
	"crypto/tls"
	"github.com/miekg/dns"
	"net"
	"os"
	"testing"
)

func TestMakeClient(t *testing.T) {
	err := os.Rename("testdata/certs/client.pem", "testdata/certs/client.back.pem")
	if err != nil {
		t.Error("pem does not exits")
	}
	_, err = MakeClient()
	if err == nil {
		t.Error("we could make client without cert")
	}
	err = os.Rename("testdata/certs/client.back.pem", "testdata/certs/client.pem")
	if err != nil {
		t.Error("we lost our certs")
	}
	_, err = MakeClient()
	if err != nil {
		t.Error("make client error", err)
	}
}

var gls net.Listener

func TestGetLisener(t *testing.T) {
	_, err := GetListener("badcerts", "badcerts", "127.0.0.1:1853")
	if err == nil {
		t.Error("we could get ls without cert")
	}
	_, err = GetListener("testdata/certs/full.pem", "testdata/certs/priv.pem", "x")
	if err == nil {
		t.Error("we could get ls without port")
	}
	ls, err := GetListener("testdata/certs/full.pem", "testdata/certs/priv.pem", ":1853")
	if err != nil {
		t.Error("get ls err", err)
	}
	gls = ls
}

func TestRun(t *testing.T) {
	defer gls.Close()
	cExit := make(chan bool)
	go Run(gls, cExit)
	// copied from client_test.go // dns lib
	m := new(dns.Msg)
	m.SetQuestion("miek.nl.", dns.TypeSOA)
	c := new(dns.Client)
	c.Net = "tcp-tls"
	c.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	r, _, err := c.Exchange(m, "127.0.0.1:1853")
	if err != nil {
		t.Fatalf("failed to exchange: %v", err)
	}
	if r == nil {
		t.Fatal("response is nil")
	}
	close(cExit)
}
