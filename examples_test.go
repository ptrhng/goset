package goset_test

import (
	"fmt"
	"sort"

	"github.com/yopeterhwang/goset"
)

func ExampleSet_Contains() {
	fruits := []string{"Apple", "Banana", "Apple", "Strawberry"}
	set := goset.From(fruits)
	fruits = set.Slice()
	sort.Strings(fruits)

	fmt.Println(fruits)

	// output:
	// [Apple Banana Strawberry]
}

func ExampleSet_Range() {
	set := goset.From([]int{1, 2, 3, 4, 5})
	set.Range(func(n int) bool {
		fmt.Println(n)
		return false
	})

	// unordered output:
	// 1
	// 2
	// 3
	// 4
	// 5
}
