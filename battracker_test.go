package battrack

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

type battrack struct {
	out io.Writer
}

func (bt *battrack) Trace(a ...interface{}) {
	bt.out.Write([]byte(fmt.Sprint(a...)))
	bt.out.Write([]byte("\n"))
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
