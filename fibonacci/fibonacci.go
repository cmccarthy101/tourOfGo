package fibonacci


// fibonacci is a function that returns
// a function that returns an int.

func fib_helper() func() int {

	prev := -1
	fib := 0

	return func() int {

		if prev == -1 {
			prev = 1
			return 0
		} else {
			tmp := fib
			fib = fib + prev
			prev = tmp
			return fib
		}
	}

}

func Fibonacci (x int) []int {

	f:= fib_helper()

	sli := make([]int, x)

	for i:=0; i < x;i++ {

		sli[i] = f()

	}

	return sli

}
