package battrack

import (
	"fmt"
	"io"
)

type battrack struct {
	out io.Writer
}

type BatTracker interface {
	Trace(...interface{})
}

type nilTracker struct{}

func New(w io.Writer) BatTracker {
	return &battrack{out: w}
}

func (bt *battrack) Trace(a ...interface{}) {
	bt.out.Write([]byte(fmt.Sprint(a...)))
	bt.out.Write([]byte("\n"))
}

func (t *nilTracker) Trace(a ...interface{}) {}

func Off() BatTracker {
	return &nilTracker{}
}
