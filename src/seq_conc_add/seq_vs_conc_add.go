package seqconcadd

import "fmt"

func SequentialVsConcurrentAddition() {
	numbers := []int{
		1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	}

	// use the pub Sum func to perform seq addition
	summedSequential := Sum(numbers)
	fmt.Println("summed sequential", summedSequential)

	// use the same func to perform concurrent addition
	// make the Sum func itself independent of seq/conc execution mode by
	// wrapping it in a closure and using a chan to pass data between roroutines
	ch := make(chan int)
	go func(sublist []int, ch chan int) {
		ch <- Sum(sublist)
	}(numbers[:len(numbers)/2], ch)

	go func(sublist []int, ch chan int) {
		ch <- Sum(sublist)
	}(numbers[len(numbers)/2:], ch)

	summedConcurrent := <-ch + <-ch
	close(ch)

	fmt.Println("summed concurrent", summedConcurrent)
}

// make this a pure func
func Sum(vals []int) (sum int) {
	for _, v := range vals {
		sum += v
	}
	return
}
