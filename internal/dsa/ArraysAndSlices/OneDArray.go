package arraysandlinkedlists

import (
	"flag"

	"github.com/golang/glog"
)

func OneDArray[V any](inputArray []V) {

	flag.Set("logtostderr", "1")
	flag.Parse()

	glog.Infoln("Arrays and Linked Lists .....")

	// one dimensional array

	oneDArray := make([]int, 10)

	glog.Infoln("printing default values of an int array...")
	// printing default values of an int array
	for _, elem := range oneDArray {
		glog.Info(elem, "")
	}

	// assigning values to an array
	for i := 0; i < len(oneDArray); i++ {
		oneDArray[i] = i
	}

	glog.Infoln("printing updated array element values...")
	// printing array values
	for _, elem := range oneDArray {
		glog.Info(elem, " ")
	}

	oneDArray1 := [4]int{1, 2, 3, 4}
	glog.Infoln("printing single line defined arrays...")
	for _, elem := range oneDArray1 {
		glog.Info(elem, " ")
	}

	glog.Infoln("printing another array....")
	var oneDArray2 [4]int
	oneDArray2 = [4]int{1, 2, 3, 4}

	for _, elem := range oneDArray2 {
		glog.Info(elem)
	}

	// Slices

	var slice1 []int

	// slice with only length
	slice1 = make([]int, 3)
	glog.Infof("slice length: %d, capacity: %d", len(slice1), cap(slice1))

	slice1[0] = 0
	glog.Infof("slice length: %d, capacity: %d", len(slice1), cap(slice1))

	slice1[1] = 1
	glog.Infof("slice length: %d, capacity: %d", len(slice1), cap(slice1))

	slice1[2] = 2
	glog.Infof("slice length: %d, capacity: %d", len(slice1), cap(slice1))

	/*
		      Append will append the element at the last of the slice.

			  If we are not defining capacity or the capacity then we use append
			  it increases the capacity to twice and it value will be appended
			  from the appended position

			  As and when we append the element the length of the slice increase
			  by one.

			  We can only access slice value using index till the length of the slice
			  if we try access value or insert value beyond length then it will
			  throw an error.
	*/

	slice1 = append(slice1, 3)
	glog.Infof("slice length: %d, capacity: %d", len(slice1), cap(slice1))

	slice1[3] = 4
	glog.Infof("slice length: %d, capacity: %d", len(slice1), cap(slice1))

	slice1 = append(slice1, 5)
	glog.Infof("slice length: %d, capacity: %d", len(slice1), cap(slice1))

	slice1 = append(slice1, 6)
	glog.Infof("slice length: %d, capacity: %d", len(slice1), cap(slice1))

	slice1 = append(slice1, 7)
	glog.Infof("slice length: %d, capacity: %d", len(slice1), cap(slice1))

	for _, elem := range slice1 {
		glog.Info(elem)
	}

	/*
	  When we define capacity then while appending, the capacity does not increae
	  till it reaches the capacity value, this give su the better performace
	*/

	// Slice2 with both elngth and capacity defined
	slice2 := make([]int, 3, 6)
	glog.Infof("slice2 length: %d, capacity: %d", len(slice2), cap(slice2))

	slice2 = append(slice2, 3)
	glog.Infof("slice2 length: %d, capacity: %d", len(slice2), cap(slice2))

	slice2[0] = 0
	glog.Infof("slice2 length: %d, capacity: %d", len(slice2), cap(slice2))

	slice2[1] = 1
	glog.Infof("slice2 length: %d, capacity: %d", len(slice2), cap(slice2))

	slice2[2] = 2
	glog.Infof("slice2 length: %d, capacity: %d", len(slice2), cap(slice2))

	slice2 = append(slice2, 3)
	glog.Infof("slice2 length: %d, capacity: %d", len(slice2), cap(slice2))

	for _, elem := range slice2 {
		glog.Info(elem)
	}

	// twoD Array

	var twoDArray [][]int

	// defining rows
	twoDArray = make([][]int, 5)

	// defining columns
	for i := 0; i < len(twoDArray); i++ {
		twoDArray[i] = make([]int, 6)
	}

	// inserting values into twoDArray....
	for i := 0; i < len(twoDArray); i++ {
		for j := 0; j < len(twoDArray[i]); j++ {
			twoDArray[i][j] = i
		}
	}

	for i := 0; i < len(twoDArray); i++ {
		for j := 0; j < len(twoDArray[i]); j++ {
			glog.Infof("twoDArray i:%d j:%d value:%d", i, j, twoDArray[i][j])
		}
	}

}
