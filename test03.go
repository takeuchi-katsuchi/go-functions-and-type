package main

func main() {
	n, m := swap(10, 20)
	println(n, m)
}

func swap(ns int, ms int) (int, int) {
	return ms, ns
}
