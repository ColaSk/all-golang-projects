package basis

import "fmt"

func Map() {
	fmt.Println("**********数组基础篇-map**************")
	// 定义map并初始化
	myMap := map[int]int{1: 1}

	for k, v := range myMap {
		fmt.Printf("map data key:%d, value: %d\n", k, v)
	}

	// 定义一个map,并分配空间
	myMap1 := make(map[int]int, 3)

	// go map是基本数据结构是hash数组+桶内的key-value数组+溢出的桶链表
	//当hash表超过阈值需要扩容增长时，会分配一个新的数组，新数组的大小一般是旧数组的2倍

	fmt.Printf("map的的长度为%d\n", len(myMap1))

	for i := 0; i < 20; i++ {
		myMap1[i] = i
	}

	fmt.Printf("map的的长度为%d\n", len(myMap1))

	for k, v := range myMap1 {
		fmt.Printf("map1 data key:%d, value: %d\n", k, v)
	}

}
