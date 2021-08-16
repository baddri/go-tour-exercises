package fibonacci

// fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func() int {
	prev := 0
	cur := 0
	return func() int {
		if cur == 0 {
			cur = 1
			return 0
		}
		temp := cur
		cur += prev
		prev = temp
		return cur
	}
}
