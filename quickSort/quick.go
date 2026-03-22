package quickSort

// 分区函数,填坑法
func partition(arr []int, left, right int) int {

	//将最左边的数当做基准
	pivot := arr[left]
	for left < right { //没有撞到，继续工作
		//先填左边的坑，找比基准小的
		for left < right && pivot <= arr[right] {
			right--
		}
		arr[left] = arr[right]
		//填右边的，找比基准大的
		for left < right && pivot >= arr[left] {
			left++
		}
		arr[right] = arr[left]
	}
	//两个指针相遇，将基准放入位置
	arr[left] = pivot //两个指针随便哪一个都行
	//返回基准位置，后面还要递归
	return left
}

func quickSort(arr []int, left, right int) {
	//递归终止条件,区间只有一个元素
	if left >= right {
		return
	}
	//拆分区间
	pinot := partition(arr, left, right)
	//递归左边
	quickSort(arr, left, pinot-1)
	//递归右边
	quickSort(arr, pinot+1, right)
}

func QuickSort(arr []int) {
	n := len(arr)
	if n <= 1 { //如果数组的长度是1/0，就不用排序
		return
	}
	quickSort(arr, 0, n-1)
}
