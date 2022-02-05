package main

func main() {
	n, m := 10, 20
	swap2(&n, &m)
	println(n, m)
}

func swap2(ns *int, ms *int) {
	*ns, *ms = *ms, *ns
}
