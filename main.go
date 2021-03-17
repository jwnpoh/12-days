// Using a for-loop to calculate the total number of gifts received over n days of Christmas.
package main

import (
	"fmt"
	"math/big"
	"strconv"
)

const usage = "How many days of Christmas is your true love gifting you gifts for?"

func main() {
	// get number of days with user input with fmt.Scan
	var input string

	fmt.Println(usage)
	fmt.Scan(&input)

	// make sure that only integers input
	d, err := strconv.ParseUint(input, 10, 64)

	if d < 0 || err != nil {
		fmt.Println("This makes no sense.")
		return
	}

	e := big.NewInt(int64(d))

	sum, total := big.NewInt(0), big.NewInt(0)

	for i := big.NewInt(1); i.Cmp(e) != 1; {
		sum.Add(sum, i)
		total.Add(total, sum)

		i.Add(i, big.NewInt(int64(1)))
	}

	fmt.Printf("The total number of gifts you would receive on the last day is %v.\n", sum)

	if d == 1 {
		fmt.Printf("The sum total number of gifts received over %d day is %v.\n", d, total)
	} else {
		fmt.Printf("The sum total number of gifts received over %d days is %v.\n", d, total)
	}
}
