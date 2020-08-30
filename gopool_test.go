package gopool

import (
	"testing"
)

func TestGopoolDemo(t *testing.T) {
	p := New(5)

	for i := 0; i < 100; i++ {
		j := i
		p.Do(func() { t.Log(j) })
	}

	p.Done()
}
