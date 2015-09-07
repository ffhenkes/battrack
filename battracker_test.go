package battrack

import (
	"bytes"
	"io"
	"testing"
)

type battrack struct {
	out io.Writer
}

func New(w io.Writer) BatTracker {
	return &battrack{out: w}
}

func TestNew(t *testing.T) {

	var buf bytes.Buffer
	battracker := New(&buf)

	if battracker == nil {
		t.Error("This should not be nil!!")
	} else {
		battracker.Trace("Hello Bat tracker!!")
		if buf.String() != "Hello Bat tracker!!\n" {
			t.Errorf("Trace should not write '%s'", buf.String)
		}
	}
}
