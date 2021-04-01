package goleak_demo

import (
	"go.uber.org/goleak"
	"testing"
)

// TODO

func chLeak() {
	ch := make(chan struct{})

	go func() {
		ch <- struct{}{}
	}()
}

func TestLeakWithGoleak(t *testing.T) {

	defer goleak.VerifyNone(t)
	chLeak()
}
