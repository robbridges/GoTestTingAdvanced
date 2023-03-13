package main

import (
	"GoTestingAdvanced/sub"
	"fmt"
)

func main() {
	result, err := sub.GitStatus()
	if err != nil {
		fmt.Errorf("gitstatus error = %s", err)
	}
	fmt.Println(result)
}
