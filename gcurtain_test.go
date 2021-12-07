package gcurtain

import (
	"strconv"
	"testing"
)

const uri = "localhost:6379"

var g = new(GCurtain)

func TestFeatureInvalidForUser(t *testing.T) {
	g.Init(uri)
	returnReceived := g.IsOpen("send_pld_producer_queue", "TEST_INVALID")
	if false != returnReceived {
		t.Errorf("Type received is different from expected! expected %s and received %s",
			"false",
			strconv.FormatBool(returnReceived))
	}
}

func TestFeatureValidForUser(t *testing.T) {
	g.Init(uri)
	returnReceived := g.IsOpen("send_pld_producer_queue", "MPA-112233")
	if true != returnReceived {
		t.Errorf("Type received is different from expected! expected %s and received %s",
			"true",
			strconv.FormatBool(returnReceived))
	}
}
