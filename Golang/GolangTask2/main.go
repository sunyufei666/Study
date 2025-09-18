package main

import "fmt"

func main() {
	/************************** 指针 **************************/
	x := 10
	modifyValue(&x)
	fmt.Println(x)

}

func modifyValue(x *int) {
	*x += 10
}
