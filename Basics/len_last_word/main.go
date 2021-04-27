package main

import "fmt"

// rootVIII 0 ms, faster than 100.00% of Go online submissions

func lengthOfLastWord(s string) int {
	var wordlen int = 0
	var collecting bool = false

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' && !collecting {
			collecting = true
		}

		if collecting {
			if s[i] == ' ' {
				break
			} else {
				wordlen++
			}
		}

	}

	return wordlen
}

func main() {
	fmt.Println(lengthOfLastWord("Hellow World"))
	fmt.Println(lengthOfLastWord(" "))
	fmt.Println(lengthOfLastWord("a "))
}
