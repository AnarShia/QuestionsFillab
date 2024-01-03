package main

import "fmt"

func main() {

	input := []string{"apple", "pie", "apple", "red", "red", "red"}

	mostFrequentWord(input)

}

func mostFrequentWord(input []string) {

	// create a map to store the number of occurences of each word
	m := make(map[string]int)

	// count the number of occurences of each word
	for _, v := range input {
		m[v]++
		// to check the map
		//fmt.Println(m)
	}

	// find the word with the highest number of occurences
	max := 0
	var maxWord string
	for k, v := range m {
		if v > max {
			max = v
			maxWord = k
		}
	}
	fmt.Println(maxWord)
}
