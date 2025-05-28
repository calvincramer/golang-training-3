package main

import (
	"time"
)

// TODO: write a function that takes the square root of a number.
// If the number is negative, then panic (we don't know what imaginary numbers are).
func DangerousSqrt(n float64) float64 {
	return 0
}

// TODO: write a function that uses DangerousSqrt() to accomplish a safe square root. In the case
// that DangerousSqrt() panics, recover and try DangerousSqrt() again with the absolute value of n.
func BubbleWrapSqrt(n float64) (ans float64) {
	return
}

// TODO: run the `GetGoogleWebpage()` function in the background using goroutines. Do NOT wait for
// the goroutine to finish. Do NOT modify the `GetGoogleWebpage()` function.
func DoBackground() {
	GetGoogleWebpage() // Make me a background task!
}

// TODO: the follow method attempts to modify the `sum` global variable using multiple goroutines to
// exploit the systems multiple CPU cores. It incorrectly calculates the sum of number from 0 to one
// million. Fix it so it is correctly calculated.
// Hint: the problem is a race condition
// Hint: you can use a lock or channels
var sum uint64 = 0

func CalcSumOneMillion() uint64 {
	worker := func(low uint64, high uint64) {
		var tempSum uint64 = 0
		for n := low; n <= high; n++ {
			tempSum += n
		}
		sum += tempSum // update global
	}
	sum = 0 // reset global
	var chunk uint64 = 100
	const end uint64 = 1_000_000
	var n uint64 = 0
	for n <= end {
		high := min(n+chunk-1, end)
		go worker(n, high)
		n += chunk
	}
	time.Sleep(time.Millisecond * 200) // bad way to wait for goroutines to finish!!!
	return sum
}

// TODO: spawn a goroutine for each number to check if it is prime or not. This method will be
// called with a large amount of very large numbers. The starting implementation calls IsPrime()
// blocking on the current thread.
//
// Hint: use IsPrime() which is defined in util.go
// Hint: use sync.WaitGroup to wait for multiple goroutines
// Hint: you should see all cores of your CPU be saturated for a few seconds. Use 'top', 'htop', or
// your system monitor to see. You should **NOT** see a single core at 100%.
//
// Note: this function should take around 1 second to compute. Adjust `REPEAT_TIMES` in
// concurrency_test.go if your machine is slower or faster.
// Note: there is no synchronization needed to modify the result slice because no two goroutines
// will modify the same index.
func IsPrimeMultiple(numbers []int64) []bool {
	var result []bool = make([]bool, len(numbers))
	for idx, num := range numbers {
		result[idx] = IsPrime(num)
	}
	return result
}

// TODO: run the `IsPrimeGoroutine()` as a goroutine and collect and return the response from the
// result channel.
func IsPrimeBackground(n int64) bool {
	return false
}

// TODO: use an unbuffered channel to "ping pong" a piece of data between two goroutines, and then
// return the response. Here's the order of operations: (Sally and Carl identify the two threads)
//
//  1. Sally: multiple number by 5
//  2. Carl: if the number is negative then take its absolute value and add 100
//  3. Sally: square the number
//  4. Carl: divide the number by it's least significant digit (skip if the digit is 0)
//  5. Sally: subtract 7
//  6. Carl OR Sally: return the number to the caller.
//
// Note: in between each operation the number is sent to the other thread.
// Hint: you can use the current goroutine as one of the goroutines. No need to create two
// additional routines.
func PingPongCalc(num int) int {
	return 0
}

// TODO: Return a buffered channel with capacity 100. Your task is to write 500 `int`s and `500`
// strings to the channel. The integers and strings may be arbitrary. After the 500 integers and
// strings, send the value `Sentinel` (in util.go) on the channel.
// You may send two `Sentinel`s, and only one is required to be after all integers and strings.
// If a `Sentinel` is sent before 500 integers and strings, it must be the case that at least 500
// integers OR 500 strings.
// Note: you will need to write the values to the channel in an asynchronous manner, since trying to
// write more than 100 values to the channel will cause the sender to block.
// Note: if this test case is taking more than a few milliseconds then it is wrong.
func MultiWriter() chan any {
	ch := make(chan any, 100)
	return ch
}

// TODO: implement a function which takes an arbitrary amount of commands on the `commands` channel,
// and stops once there is *any* value received on the `done` channel.
// `Command` is a struct which has a `Operation` enum and a slice of arguments.
// This function needs to keep track of a number, and update it according to the commands, and then
// return it at the end of execution. The number should start at 0.
//
// For example, starting with 0:
//
//  1. Command(OP_SET 50) -> 50
//  2. Command(OP_ADD 1 7 20) -> 50 + 1 + 7 + 20 = 78
//  3. Command(OP_MULT -1) -> 78 * -1 = -78
//  4. Message received on `done`, -78 returned.
//
// Note: use the `select` feature.
// Note: addition arguments to operations that require some or none may safely be ignored
// Note: you may assume that all input commands are well-formed and have their required arguments
func CommandConsumer(commands chan Command, done chan bool) int {
	return 0
}

// TODO: use tools from the 'sync' package to fix the following code below.
// Do not change the following aspects of PerformAccounting():
//
//  1. PerformAccounting first clears the CategoryTotals global
//  2. PerformAccounting spawns three Accountants and splits the transactions equally among them
//  3. PerformAccounting waits for all goroutine Accountants to finish before returning
//
// Note: you may *minimally* change TestPerformAccounting() if using a different data type for
// CategoryTotals.
var CategoryTotals map[string]int = map[string]int{}

func Accountant(transactions []Transaction, done chan<- bool) {
	for _, trans := range transactions {
		curTotal, exists := CategoryTotals[trans.Category]
		if !exists {
			CategoryTotals[trans.Category] = trans.Amount
		} else {
			CategoryTotals[trans.Category] = curTotal + trans.Amount
		}
	}
	done <- true
}

func PerformAccounting(transactions []Transaction) {
	clear(CategoryTotals) // Do not remove this line

	// Share tasks among multiple accountants
	third := len(transactions) / 3
	twoThird := (len(transactions) * 2) / 3

	doneCh := make(chan bool, 3)
	go Accountant(transactions[:third], doneCh)
	go Accountant(transactions[third:twoThird], doneCh)
	go Accountant(transactions[twoThird:], doneCh)

	// Wait for accountants
	doneCount := 0
	for {
		for range doneCh {
			doneCount += 1
			if doneCount >= 3 {
				return
			}
		}
	}
}
