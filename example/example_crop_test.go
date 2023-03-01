package example_test

import (
	"GoTestingAdvanced/example"
	"fmt"
	_ "image/png"
	"io"
)

var file string = "This is not used"

func ExampleCrop() {

	var r io.Reader

	img, err := example.Decode(r)
	if err != nil {
		panic(err)
	}

	err = example.Crop(img, 0, 0, 20, 20)
	if err != nil {
		panic(err)
	}
	var w io.Writer
	if err != nil {
		panic(err)
	}

	err = example.Encode(img, w)
	if err != nil {
		panic(err)
	}

	fmt.Println("See out.jpg for the cropped image")

	// Output:
	// See out.jpg for the cropped image

}
