package trace

import (
	"io"
	"fmt"
)

type Tracer interface {
	// Tracer is the interface that describes an object capable of 
	// tracing events throughout code
	// arguments:
	// * ...interface{ : takes in zero or more arguments of any type
	Trace(...interface{})
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}