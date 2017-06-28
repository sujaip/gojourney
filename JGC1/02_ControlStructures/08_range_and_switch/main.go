package main

import "fmt"

func main() {
	atoz := "the quick brown fox jumps over the lazy dog"

	vowels, consonants, zeds := 0, 0, 0

	for _, r := range atoz {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		case 'z':
			zeds++
			/* Here we explicitly fallthrough as zeds are consonants */
			fallthrough
		default:
			consonants++
		}
	}

	fmt.Printf("Vowels: %d; Consonants: %d\n Zeds: %d", vowels, consonants, zeds)
}
