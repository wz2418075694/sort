package shellSort

//思路：让选择排序的外层套一个逐渐减少步长的循环

func ShellSort(arr []int) {

	//1.设置步长，初始为数组长度的一半，然后逐渐减半
	n := len(arr)
	for gap := n / 2; gap > 0; gap = gap / 2 {
		//从gap开始，每一个都是自己那组未排序区的第一个元素
		for i := gap; i < n; i++ {
			key := arr[i] //哨兵
			//j是已排序区的最后一个元素
			j := i - gap
			for ; j >= 0 && arr[j] > key; j = j - gap {
				//后移
				arr[j+gap] = arr[j]
			}
			//插入，j此时指向的是已排序的元素，所以要插入到它的后面
			arr[j+gap] = key
		}
	}
}
