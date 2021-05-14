/*
Given a string s and an integer k, return the length of the longest substring
of s such that the frequency of each character in this substring is greater
than or equal to k.
*/

package main

import "fmt"

func calcLongestSubstring(str string, k int) int {
	if k == 1 {
		return len(str)
	}
	if k > len(str) {
		return 0
	}

	uniqChar := getUniqChar(str)
	maxLength := 0

	for currentUniq := 1; currentUniq <= uniqChar; currentUniq++ {
		mp := make(map[rune]int)
		uniq, numofK, end, start := 0, 0, 0, 0
		for end < len(str) {
			if uniq <= currentUniq {
				ele := str[end]
				if mp[rune(ele-'a')] == 0 {
					uniq++
				}
				mp[rune(ele-'a')]++
				if mp[rune(ele-'a')] == k {
					numofK++
				}
				end++
			} else {
				ele := str[start]
				if mp[rune(ele-'a')] == k {
					numofK--
				}
				mp[rune(ele-'a')]--
				if mp[rune(ele-'a')] == 0 {
					uniq--
				}
				start++
			}
			// if number of uniq element is equivalent to given current Uniq element
			// and these current uniq element have atleast k characters
			if uniq == currentUniq && uniq == numofK {
				maxLength = max(maxLength, end-start)
			}
		}
	}
	return maxLength
}

func max(first int, rest ...int) (max int) {
	max = first
	for _, value := range rest {
		if value > max {
			max = value
		}
	}
	return
}

func getUniqChar(str string) int {
	chars := make([]bool, 26)
	for _, ch := range str {
		chars[ch-'a'] = true
	}
	count := 0
	for _, val := range chars {
		if val == true {
			count++
		}
	}
	return count
}

func main() {
	str := "xyxyyz"
	k := 2
	res := calcLongestSubstring(str, k)
	fmt.Printf("length of longest substring is %d for %s with atleast %d repeating characters\n", res, str, k)
}
