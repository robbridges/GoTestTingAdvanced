package comparison

import (
	"reflect"
	"testing"
)

func TestSquare(t *testing.T) {
	arg := 2
	want := 4
	got := Square(arg)
	if got != want {
		t.Errorf("Square(%d) We got %d, but wanted %d", arg, got, want)
	}

}

func TestDog(t *testing.T) {
	morty := Dog{
		Name: "Morty",
		Age:  3,
	}
	morty2 := Dog{
		Name: "Morty",
		Age:  3,
	}
	if morty != morty2 {
		t.Errorf("morty != morty2")
	}

	odie := Dog{
		Name: "Odie",
		Age:  12,
	}

	if morty == odie {
		t.Errorf("Mory == odie")
	}
}

func TestPtr(t *testing.T) {
	morty := &Dog{
		Name: "Morty",
		Age:  3,
	}
	morty2 := &Dog{
		Name: "Morty",
		Age:  3,
	}
	t.Logf("morty=%p, morty2=%p", morty, morty2)
	if morty == morty2 {
		t.Errorf("morty == morty2")
	}
}

func TestDogWithFn(t *testing.T) {
	morty := DogWithFn{
		Name: "Morty",
		Age:  3,
		Fn: func() {

		},
	}
	morty2 := DogWithFn{
		Name: "Morty",
		Age:  3,
		Fn: func() {

		},
	}
	if reflect.DeepEqual(morty, morty2) {
		t.Errorf("morty != morty2")
	}
}
