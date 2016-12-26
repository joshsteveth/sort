package sort

import (
	"math/rand"
	"os"
	"testing"
)

var numElement int

func TestMain(m *testing.M) {
	numElement = 200

	os.Exit(m.Run())
}

func createNewSlice(n int) []int {
	var s []int
	//create random int as many as n
	for i := 0; i < n; i++ {
		s = append(s, rand.Intn(n))
	}

	return s
}
