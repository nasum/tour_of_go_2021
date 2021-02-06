package main

import "golang.org/x/tour/tree"

func Walk(t *tree.Tree, ch chan int) {
	var walk_sub func(*tree.Tree)
	walk_sub = func(t *tree.Tree) {
		if t != nil {
			walk_sub(t.Left)
			ch <- t.Value
			walk_sub(t.Right)
		}
	}
	walk_sub(t)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v1 := range ch1 {
		v2, ok := <-ch2
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		println("tree:", v)
	}
	println("same 1,1", Same(tree.New(1), tree.New(1)))
	println("same 1,2", Same(tree.New(1), tree.New(2)))
}
