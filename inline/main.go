package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	z := max(1, 2)
	println(z)
}
