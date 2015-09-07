package battrack

import "fmt"

type BatTracker interface {
	Trace(...interface{})
}

func (bt *battrack) Trace(a ...interface{}) {
	bt.out.Write([]byte(fmt.Sprint(a...)))
	bt.out.Write([]byte("\n"))
}
