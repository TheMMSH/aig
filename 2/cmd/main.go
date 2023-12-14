package main

import (
	second "aig/2"
	"fmt"
)

func main() {
	fmt.Printf("abccccccda -> \"%s\"\n", second.Rearrange("abccccccda"))
	fmt.Printf("aa -> \"%s\"\n", second.Rearrange("aa"))
	fmt.Printf("aaaaabbbbbcccc -> \"%s\"\n", second.Rearrange("aaaaabbbbbcccc"))
	fmt.Printf("zzxxccaaww -> \"%s\"\n", second.Rearrange("zzxxccaaww"))
}
