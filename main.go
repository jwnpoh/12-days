// Using a for-loop to calculate the total number of gifts received over n days of Christmas.
package main

import (
	"fmt"
	"os"
	"strconv"
)

// If no arguments given.
const usage = "How many days of Christmas is your true love gifting you gifts on?"

func main() {
	arg := os.Args

	if len(arg) != 2 {
		fmt.Println(usage)
		return
	}

	// get number of days
	d, err := strconv.Atoi(arg[1])

	if err != nil {
		fmt.Println(usage + " (in numbers)")
		return
	}

	var sum, total int

	for i := 1; i <= d; i++ {
		sum += i
		total += sum
	}
	fmt.Printf("The total number of gifts you would receive on the last day is %d.\n", sum)

	if d == 1 {
		fmt.Printf("The sum total number of gifts received over %d day is %d.\n", d, total)
	} else {
		fmt.Printf("The sum total number of gifts received over %d days is %d.\n", d, total)
	}
}
