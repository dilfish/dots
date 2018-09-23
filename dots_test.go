package dots

import (
	"testing"
)

func TestMakeClient(t *testing.T) {
	_, err := MakeClient()
	if err != nil {
		t.Error("make client error", err)
	}
}
