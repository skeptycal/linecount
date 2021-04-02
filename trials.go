package linecount

import (
	"bytes"
	"fmt"
)

const (
	defaultTrialCount = 10
	defaultSampleSize = 32768
)

type Content interface{}

func (c Content) String() string {
	return fmt.Sprintf("%v", c)
}

type Trial interface {
	String() string
	Count() int
	Result() interface{}
}

type trial struct {
	n           int
	content     interface{}
	keepContent bool
	f           func() int
}

func (t *trial) Count() int {
	var counter int = 0
	for _, c := range t.content {
		if c == '\n' {
			counter += 1
		}
	}
	return counter
}

func (t *trial) run() interface{} {
	buf := bytes.Buffer{}
	for i := 0; i < t.n; i++ {
		buf.WriteByte(randByte())
	}
	return buf.Bytes()
}

func (t *trial) Result() interface{} {
	if t.content == nil {
		t.content = t.run()
	}
	return t.content
}
