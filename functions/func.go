package functions

import "fmt"

func adder() func(int) int {
	//fmt.Println("Adder Function accessed.")
	sum := 0
	return func(x int) int {
		//fmt.Printf("x is equal to %d\n", x)
		//fmt.Println("Inside func accessed.")
		sum += x
		return sum
	}
}

func posneg(){

	//fmt.Println("Starts HERE")
	pos, neg := adder(), adder()

	//fmt.Println("i loop starts HERE")
	for i := 0; i < 10; i++ {

		//fmt.Printf("Iteration %d:\n:", i)
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
