package parallel

import (
	"fmt"
	"testing"
)

//func TestSomething(t *testing.T) {
//	t.Parallel()
//	time.Sleep(time.Second)
//}
//
//func TestA(t *testing.T) {
//	t.Parallel()
//	time.Sleep(time.Second)
//}

//func TestB(t *testing.T) {
//	fmt.Println("setup")
//	defer fmt.Println("deferred teardown")
//	t.Run("group", func(t *testing.T) {
//
//		t.Run("sub1", func(t *testing.T) {
//			t.Parallel()
//			//run test
//			time.Sleep(time.Second)
//			fmt.Println("Sub one done")
//		})
//		t.Run("sub2", func(t *testing.T) {
//			t.Parallel()
//			// run test2
//			time.Sleep(time.Second)
//			fmt.Println("Sub two done")
//		})
//	})
//	fmt.Println("teardown")
//}

func TestGotcha(t *testing.T) {
	testCases := []struct {
		arg  int
		want int
	}{
		{2, 4},
		{3, 9},
		{4, 16},
	}

	for _, tc := range testCases {
		// we need to shadow this variable so that the parellel tests run it with each one.
		// Please do not delete if deleted you will get false positives on any failures
		tc := tc
		t.Run(fmt.Sprintf("arg=%d", tc.arg), func(t *testing.T) {
			t.Parallel()
			t.Logf("tcarg: %d, tc:want %d", tc.arg, tc.want)
			if tc.arg*tc.arg != tc.want {
				t.Errorf("%d^ != %d", tc.arg, tc.want)
			}
		})
	}
	//for i := 0; i < 10; i++ {
	//	t.Run(fmt.Sprintf("i=%d", i), func(t *testing.T) {
	//		t.Parallel()
	//		t.Logf("Testing with i=%d", i)
	//	})
	//}
}
