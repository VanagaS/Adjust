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

	for len(mapp) != size {
		// get a random integer within given size (10, here)
		k := rand.Intn(size + 1)

		// we do not want 0 printed
		if k == 0 {
			continue
		}

		// check if this random generated number is already printed
		if !mapp[k] {
			// print the number
			print(k, "\n")
		}

		// save the number
		mapp[k] = true
	}
}
