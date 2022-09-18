package basis

import "fmt"

func List() {
	println("********数组基础篇-固定大小数字定义*********")
	// 定义固定长度数组 默认值初始化为[0,0,0...]
	var myArr [5]int

	// 初始化一个数组
	myArr = [5]int{1, 2}

	// 初始化指定元素

	myArr = [5]int{2: 1, 3: 2}

	for i := 0; i < len(myArr); i++ {
		fmt.Printf("固定长度数组第%d个值<%d>\n", i, myArr[i])
	}

	// 定义固定长度数组，默认值初始化为[1,2,3,0,0...]
	myArr1 := [5]int{1, 2, 3}

	for i := 0; i < len(myArr1); i++ {
		fmt.Printf("固定长度数组默认初始化第%d个值<%d>\n", i, myArr1[i])
	}

}

func MultiList() {
	fmt.Println("**********数组基础篇-多维数组**************")
	//声明多维数组
	var myArr [4][2]int
	// 使用数组字面量来声明并初始化一个二维数组
	myArr = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	// 使用索引初始化二维数组
	myArr = [4][2]int{1: {20, 21}, 3: {40, 41}}
	for i := 0; i < len(myArr); i++ {
		for j := 0; j < len(myArr[i]); j++ {
			fmt.Printf("多维数组第%d-%d个值<%d>\n", i, j, myArr[i][j])
		}
	}

}

func SliceList() {
	fmt.Println("**********数组基础篇-动态数组与切片**************")
	//声明一个动态数组并初始化，他就是一个切片,长度为3
	slice := []int{1, 2, 3}

	for i := 0; i < len(slice); i++ {
		fmt.Printf("动态数组(切片)第%d个值<%d>\n", i, slice[i])
	}
	//声明一个切片，未分配空间，需要使用需要分配空间
	var slice1 []int
	// 为分片分配空间
	// 切片长度与容量不同，长度表示左指针到右指针的距离，容量表示左指针到底层数组末尾的距离
	slice1 = make([]int, 2, 4)
	fmt.Printf("切片的的长度为%d,容量为%d\n", len(slice1), cap(slice1))
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("动态数组2(切片)第%d个值<%d>\n", i, slice1[i])
	}

	// 切片的追加与截取
	//切片的扩容机制，append时，如果长度增加后超过容量，则将容量增加两倍

	for i := 0; i < 5; i++ {
		slice1 = append(slice1, i)
	}
	fmt.Printf("切片的的长度为%d,容量为%d\n", len(slice1), cap(slice1))

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("动态数组2(切片)第%d个值<%d>\n", i, slice1[i])
	}

}
