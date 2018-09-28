package dots

import (
	"github.com/miekg/dns"
	"testing"
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
	_, _, err := c.Exchange(m, "127.0.0.1:853")
	return err
}

func TestGetLisener(t *testing.T) {
	ls, err := GetListener("testdata/certs/client.pem", "testdata/certs/client.key")
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
