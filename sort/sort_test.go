package sort

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

// List of algorithms to be tested
var sortAlgorithms = []func(elements []int) []int {
	BubbleSort,
	SelectionSort,
}

// Test all of the algorithms
func TestSortUnique(t *testing.T) {
	expected := []int{1, 3, 4, 5, 7}

	for _, algo := range sortAlgorithms {
		elements := []int{5, 1, 4, 3, 7}
		algo(elements)
		assert.Equal(t, expected, elements, fmt.Sprintf("Expectations do not match the actual values."), expected, elements)
	}
}

func TestSortWithDuplicates(t *testing.T) {
	expected := []int{1, 1, 1, 3, 4, 4, 5, 7}

	for _, algo := range sortAlgorithms {
		elements := []int{1, 5, 1, 4, 3, 7, 1, 4}
		algo(elements)
		assert.Equal(t, expected, elements, fmt.Sprintf("Expectations do not match the actual values."), expected, elements)
	}
}
