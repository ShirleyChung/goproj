package main

import ("fmt";
	"os")

func main() {
	fmt.Println("1314 世界!")

	var a = 100
	var b = true
	fmt.Println("a = ", a)
	if b {
		fmt.Println("it is true")
	}
	_, err := os.Open("demo.go")
	if err != nil {
		fmt.Println("file open error:", err)
	} else {
		fmt.Println("file opened!")
	}
}
