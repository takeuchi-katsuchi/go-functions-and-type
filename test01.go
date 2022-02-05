package main

func main() {
	data := []int{19, 86, 1, 12}
	var sumData int
	for i := 0; i < len(data); i++ {
		sumData += data[i]
	}
	println(sumData)
}
