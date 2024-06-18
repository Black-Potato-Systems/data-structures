package arraysandlinkedlists_test

import (
	"fmt"
	"testing"

	arraysandlinkedlists "github.com/blackpotato/data-structures/internal/dsa/ArraysAndLinkedLists"
)

func TestOneDArray(t *testing.T) {

	sampleArray := []int{1, 2, 3, 4, 5}

	arraysandlinkedlists.OneDArray[int](sampleArray)

	fmt.Printf("t: %v\n", t)
}
