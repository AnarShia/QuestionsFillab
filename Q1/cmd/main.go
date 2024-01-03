package main

import "fmt"

func main() {
	// given input
	input := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	// sort the input by the number of char 'a' in each string in descending order
	sortByCharA(input)
	// print the sorted input
	fmt.Println(input)
}

func sortByCharA(input []string) {

	// foreach string in input
	for i := 0; i < len(input); i++ {

		// foreach string in input starting from i + 1
		// compare the number of char 'a' in input[i] and input[j]
		for j := i + 1; j < len(input); j++ {

			// if the number of char 'a' equal then compare the length of input[i] and input[j]
			if numberOfCharA(input[i]) == numberOfCharA(input[j]) {

				if len(input[i]) < len(input[j]) {
					input[i], input[j] = input[j], input[i]
				}
				// else if the number of char 'a' in input[i] is less than the number of char 'a' in input[j]
			} else if numberOfCharA(input[i]) < numberOfCharA(input[j]) {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
}

func numberOfCharA(input string) int {
	count := 0
	for _, v := range input {
		if v == 'a' {
			count++
		}
	}
	return count
}
