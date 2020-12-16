// Using a for-loop to calculate the total number of gifts received over n days of Christmas.
package main

import (
	"fmt"
	"strconv"
	//"os"
)

const usage = "How many days of Christmas is your true love gifting you gifts for?"

func main() {

	// get number of days from command line arguments (os and strconv required)
	/*	arg := os.Args

			// If no arguments given.
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
	*/

	// get number of days with user input with fmt.Scan
	var u string

	fmt.Println(usage)
	fmt.Scan(&u)

	// make sure that only integers input
	d, err := strconv.Atoi(u)

	if d < 0 || err != nil {
		fmt.Println("This makes no sense.")
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
