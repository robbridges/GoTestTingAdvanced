package random

import (
	"math/rand"
	"time"
)

func Pick(numbers []int) int {
	seed := time.Now().UnixNano()
	rand := rand.New(rand.NewSource(seed))
	return numbers[rand.Intn(len(numbers))]
}
