package main

import "fmt"

func makeSet(words []string) map[string]struct{} {
	res := make(map[string]struct{})
	for _, val := range words {
		res[val] = struct{}{}
	}
	return res
}

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	set := makeSet(words)
	for key := range set {
		fmt.Print(key, " ")
	}
}
