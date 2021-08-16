package binary_tree

import (
	"testing"

	"golang.org/x/tour/tree"
)

func TestWalk(t *testing.T) {
	ch := make(chan int)
	// res := make([]int, 10)
	go func() {
		Walk(tree.New(1), ch)
		close(ch)
	}()
	for {
		_, ok := <-ch
		if !ok {
			return
		}
		//fmt.Println(v)
	}
}

func TestSame(t *testing.T) {
	t.Run(
		"Shouldn't fail when t1 == t2",
		func(t *testing.T) {
			same := Same(tree.New(1), tree.New(1))
			if !same {
				t.Fatalf("Same() expect return true but got %v", same)
			}
		},
	)
	t.Run(
		"Should fail when t1 != t2",
		func(t *testing.T) {
			same := Same(tree.New(1), tree.New(2))
			if same {
				t.Fatalf("Same() expect return false but got %v", same)
			}
		},
	)
}
