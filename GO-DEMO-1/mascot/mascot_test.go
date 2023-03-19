package mascot_test

import (
	"testing"

	"github.com/mzhou356/learn-golang/GO-DEMO-1/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong mascot :(")
	}

}
