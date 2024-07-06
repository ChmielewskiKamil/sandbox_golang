package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	v1 := make(chan int)
	v2 := make(chan int)

	go func() {
		Walk(t1, v1)
		close(v1)
	}()

	go func() {
		Walk(t2, v2)
		close(v2)
	}()

	for {
		x, ok1 := <-v1
		y, ok2 := <-v2
		if ok1 != ok2 || x != y {
			return false
		}

		if !ok1 {
			break
		}
	}

	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)

	fmt.Println(Same(t1, t2))
}
