package math

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum([]int{10, -2, 3})
	want := 11
	if got != want {
		t.Errorf("we wanted %d, but got %d", want, got)
	}
	fmt.Println("PASS")
}
