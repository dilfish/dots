package dots

import (
	"testing"
    "time"
)

func TestMakeClient(t *testing.T) {
	_, err := MakeClient()
	if err != nil {
		t.Error("make client error", err)
	}
}


func Out(c chan bool) {
    time.Sleep(time.Second)
    close(c)
}


func TestGetLisener(t *testing.T) {
    ls, err := GetListener("testdata/certs/client.pem", "testdata/certs/client.key")
    if err != nil {
        t.Error("get ls err", err)
    }
    defer ls.Close()
    cExit := make(chan bool)
    go Out(cExit)
    Run(ls, cExit)
}
