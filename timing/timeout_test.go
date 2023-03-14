package timing

import (
	"flag"
	"testing"
	"time"
)

var (
	timeoutFlag int64
)

func init() {
	flag.Int64Var(&timeoutFlag, "timeoutMultiplier", 1, "set this to a higher value value if you are "+
		"on a slower pc and experiencing failure due to unexpeted timeOuts")
}

func TestDoThing(t *testing.T) {
	timeoutMultiplier = time.Duration(timeoutFlag)
	err := DoThing()
	if err != nil {
		t.Errorf("DoThing() err = %s; want nil", err)
	}
}

func TestDoThing_inject(t *testing.T) {
	// change time After to send a time to the channel
	timeAfter = func(d time.Duration) <-chan time.Time {
		ch := make(chan time.Time)
		go func() { ch <- time.Now() }()
		return ch
	}
	// there should bne an error here, we are returning a time
	err := DoThing()
	if err == nil {
		t.Errorf("DoThing() err = nil; want %s", ErrTimeout)
	}
	// same process except not sending a current time to the channel
	timeAfter = func(d time.Duration) <-chan time.Time {
		ch := make(chan time.Time)
		go func() {}()
		return ch
	}
	// no error
	err = DoThing()
	if err != nil {
		t.Errorf("DoThing() err = %s; want nil", err)
	}

	timeAfter = time.After
}
