package main

import "fmt"

func main() {
	input := "aabbbaccccddddddddddddddddddddddddeepoilrigfencingcloud"
	longest_sub := substring_pallindrome(input)
	fmt.Println("Longest Substring: ")
	fmt.Printf("%s: %d\n\n", longest_sub, len(longest_sub))

	fmt.Println("Full Array: ")
	fmt.Println(string(input))
	fmt.Println()

	compressed := count_compress(input)
	fmt.Println("Compressed Array: ")
	fmt.Println(string(compressed))
}
