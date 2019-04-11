package main

import (
	"fmt"
)

func slice() {
	//切片是一种数据结构,这种数据结构便于使用和管理数据集合。切片是围绕动态数组的概念
	//构建的,可以按需自动增长和缩小。切片的动态增长是通过内置函数 append 来实现的

	// slice has three field are pointer, len , capacity

	// 如果在 [] 运算符里指定了一个值,那么创建的就是数组而不是切片.
	array := [3]int{1, 2, 3}
	slice := []int{2, 3, 3}
	slice = make([]int, 3)    //len and capacity == 3
	slice = make([]int, 3, 4) //len is 3,可以访问 3 个元素, 底层数组拥有 4 个元素, 剩余的 2 个元素可
	//以在后期操作中合并到切片,可以通过切片访问这些元素
	sliceStr := []string{9: "lkl"} //初始化100 elements
	fmt.Println(array)
	fmt.Println(slice)
	fmt.Printf("sliceStr is %s\n", sliceStr)
	//程序声明一个值为 nil 的切片, point = nil, len = 0, capacity = 0.
	var sliceInt []int

	c := cap(sliceInt)

	fmt.Printf("cap is %d\n", c)

}
