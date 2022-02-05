package main

import "fmt"

func main() {
	// 100をHex型として代入
	var hex Hex = 100
	// Stringメソッドを呼び出す
	fmt.Println(hex.String())
}

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}
