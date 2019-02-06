package fibonacci

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {

	x:= 10
	sliTest := Fibonacci(x)

	sliTruth:=[]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	if !reflect.DeepEqual(sliTest, sliTruth) {

		t.Errorf("expected array of %v, but got array: %v",sliTruth , sliTest)
	}

}