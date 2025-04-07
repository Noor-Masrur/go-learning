package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
	// fmt.Println(lengthOfLongestSubstring("bbbbb"))
	// fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func max(var1 int, var2 int) int {
	if var1 >= var2 {
		return var1
	}

	return var2
}

func lengthOfLongestSubstring(s string) int {
	hash := make(map[rune]int)
	start := 0
	maxx := 0
	for index, val := range s {
		idx, exist := hash[val]
		if exist {
			start = max(start, idx+1)
		}
		hash[val] = index
		fmt.Println("index:", index, "idx:", idx, "start:", start, "maxx:", maxx)
		maxx = max(maxx, index-start+1)

	}

	return maxx
}
