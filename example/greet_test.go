package example

import "fmt"

func ExampleHello() {
	result := Hello("Rob")
	fmt.Println(result)

	// Output:
	// Hello Rob
}

func ExampleDemo_DemoHello() {
	result := Hello("Rob")
	fmt.Println(result)

	// Output:
	// Hello Rob
}

func ExamplePage() {
	checkIns := map[string]bool{
		"Rob":       true,
		"Anna":      false,
		"Shawn":     false,
		"Zach":      true,
		"Some Hobo": false,
	}
	Page(checkIns)

	// Unordered Output:
	// Paging Anna; please see the front desk to check in
	// Paging Shawn; please see the front desk to check in
	// Paging Some Hobo; please see the front desk to check in
}
