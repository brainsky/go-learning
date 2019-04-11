package main

import (
	"fmt"
)

func displayArray() {

	array := [5]int{10, 13, 15, 20, 22} //initialize

	array = [...]int{20, 77, 30, 23, 10} //use ... auto caculate length

	array = [5]int{1: 11, 2: 22}

	fmt.Println(len(array))

	var strarray [5]string

	strarray2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

	strarray = strarray2

	fmt.Println(len(strarray))

	var array1 [3]*string

	array2 := [3]*string{new(string), new(string), new(string)}

	*array2[0] = "Blue"
	*array2[1] = "Red"
	*array2[2] = "Yellow"

	array1 = array2

	fmt.Println(len(array1))
}

func multiArray() {
	// 二维 数组 Two-dimensional array

	var array [4][2]int
	array = [4][2]int{{1, 0}, {22, 66}, {44, 65}, {7, 8}}
	array = [4][2]int{1: {2, 0}, 3: {3, 2}}

	fmt.Println(array)

	var array1 [2][2]int

	var array2 [2][2]int
	array2[0][0] = 1
	array2[0][1] = 2
	array2[1][0] = 3
	array2[1][1] = 4

	array1 = array2

	var array3 [2]int = array1[1]

	fmt.Println(array3)

}

func main() {
	displayArray()

	multiArray()

	slice()
}
