package main

import "fmt"

func main() {
	input := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	sortByCharA(input)
	fmt.Println(input)
}

func sortByCharA(input []string) {

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {

			if numberOfCharA(input[i]) == numberOfCharA(input[j]) {

				if len(input[i]) < len(input[j]) {
					input[i], input[j] = input[j], input[i]
				}

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
