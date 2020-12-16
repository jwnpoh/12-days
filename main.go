// Using a for-loop to calculate the total number of gifts received over n days of Christmas.
package main

import (
	"fmt"
	//"os"
	//"strconv"
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

	// get number of days with interactive prompt (os and strconv not required)
	var d int

	fmt.Println(usage)
	fmt.Scan(&d)

	var sum, total int

	if d < 0 {
		fmt.Println("This makes no sense.")
		return
	}

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
