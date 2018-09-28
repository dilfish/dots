package dots

import (
	"github.com/miekg/dns"
	"testing"
    "fmt"
)

func TestMakeClient(t *testing.T) {
	_, err := MakeClient()
	if err != nil {
		t.Error("make client error", err)
	}
}

func sendDns() error {
	m := new(dns.Msg)
	m.Id = dns.Id()
	m.RecursionDesired = true
	m.Question = make([]dns.Question, 1)
	m.Question[0] = dns.Question{"baidu.com.", dns.TypeA, dns.ClassINET}
	c := new(dns.Client)
    c.Net = "tcp-tls"
    conn, err := c.Dial("local.xn--oht.com:1853")
    if err != nil {
        println("dial", err)
        return err
    }
    err = conn.WriteMsg(m)
    if err != nil {
        println("write", err)
        return err
    }
    bt := make([]byte, 5400)
    n, err := conn.Read(bt)
    if err != nil {
        fmt.Println("read", n, err, bt)
    }
    return err
}

func TestGetLisener(t *testing.T) {
	ls, err := GetListener("testdata/certs/full.pem", "testdata/certs/priv.pem", "1853")
	if err != nil {
		t.Error("get ls err", err)
	}
	defer ls.Close()
	cExit := make(chan bool)
	go Run(ls, cExit)
	err = sendDns()
	if err != nil {
		t.Error("send dns err", err)
	}
	close(cExit)
}
