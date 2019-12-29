package main

import (
	"fmt"
	"gopl.io/ch6/exercise6.1/intset"
)

func main() {
	var x intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	fmt.Println(x.Len())    // 4

	x.Remove(42)
	fmt.Println(&x)      // "{1 9 144}"
	fmt.Println(x.Len()) // 3

	x.Remove(192)
	fmt.Println(&x)      // "{1 9 144}"
	fmt.Println(x.Len()) // 3

	fmt.Println("Copy")
	n := x.Copy()
	x.Add(222)
	fmt.Println(n)       // "{1 9 144}"
	fmt.Println(n.Len()) // 3
	fmt.Println(&x)      // "{1 9 144}"
	fmt.Println(x.Len()) // 3

	fmt.Println("Clear")
	x.Clear()
	fmt.Println(&x)      // "{1 9 144}"
	fmt.Println(x.Len()) // 3

	x.AddAll(3, 6, 9)
	fmt.Println(&x)      // "{1 9 144}"
	fmt.Println(x.Len()) // 3

	y := x.Copy()
	y.AddAll(4, 6, 8, 10, 12)
	x.AddAll(12, 15)
	y.IntersectWith(&x)
	fmt.Println(y)
}
