package main

//type Tree struct {
//    Left  *Tree
//    Value int
//    Right *Tree
//}

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func walk(t *tree.Tree, ch chan int) {
	if t != nil {
		walk(t.Left, ch)
		ch <- t.Value
		walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	var channelT1 = make(chan int)
	var channelT2 = make(chan int)

	go Walk(t1, channelT1)
	go Walk(t2, channelT2)

	for i := range channelT1 {
		if i != <-channelT2 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))

}
