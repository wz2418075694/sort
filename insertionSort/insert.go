package insertionSort

func InsertionSort(arr []int) {

	//计算数组的长度
	n := len(arr)
	//让未排序区从1开始，一个数不用遍历，没有意义，
	for i := 1; i < n; i++ {
		//设置哨兵,这就是那张要插入的牌
		key := arr[i]
		//内循环用于比较和移动
		//j初始化为已排序区的最后一个元素(也就是要和哨兵比较的那个元素)
		//j逐渐往前移动，然后挨个进行比较
		j := i - 1
		for ; j >= 0 && (arr[j] > key); j-- {
			//向后移动
			arr[j+1] = arr[j]
		}
		//找到了插入的位置，此时j指向的是比哨兵小的那个数，所以应插入到j的后一个位置
		arr[j+1] = key
	}

}
