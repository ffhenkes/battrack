package battrack

import (
	"bytes"
	"testing"
)

/* Check the trace itself */
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

/* Check disabling the trace */
func TestOff(t *testing.T) {
	var silent BatTracker = Off()
	silent.Trace("silent")
}
