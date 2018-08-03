package main

import (
	"sort"
	"fmt"
)

func main() {
	//Create a slice of int
	a := []int{3,6,4,7,9,0}
	sort.Ints(a)

	for _, v := range a{
		fmt.Println(v)
	}
}
