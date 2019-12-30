package main

import (
	"fmt"
	"gopl.io/ch7/exercise7.3/treesort"
	"math/rand"
	"time"
)

func main() {
	var tree treesort.Tree
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 50; i++ {
		tree.Add(rand.Int() % 50)
	}

	fmt.Println(&tree)
	fmt.Println(tree.String())
}
