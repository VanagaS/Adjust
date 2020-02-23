package main

import (
	"math/rand"
)

func main() {
	o_of_n_time_o_of_n_space(10)
}

func o_of_n_time_o_of_n_space(size int) {
	// create a map to contain the random numbers
	mapp := make(map[int]bool)

	// consider a case, where consecutive numbers are not printed
	previous := 0

	// Print the numbers from 1 to 10 in random order to the terminal.
	for len(mapp) != size {
		// get a random integer within given size (10, here)
		// note: default rand is deterministic
		k := rand.Intn(size + 1)

		// we do not want 0 printed nor the next and previous values of previous number
		// this would normally add an overhead.
		if k == 0 || previous == k || previous == k+1 || previous == k-1 {
			continue
		}

		// check if this random generated number is already printed
		if !mapp[k] {
			// print the number
			print(k, "\n")

			// save the number as the previous number
			previous = k
		}

		// save the number
		mapp[k] = true
	}
}
