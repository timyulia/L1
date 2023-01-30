package main

import "fmt"

func swap(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func main() {
	var a, b int
	fmt.Print("enter a ")
	fmt.Scanf("%d", &a)
	fmt.Print("enter b ")
	fmt.Scanf("%d", &b)
	fmt.Printf("before: a = %d b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("after: a = %d b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("after all: a = %d b = %d\n", a, b)
}
