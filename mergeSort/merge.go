package mergeSort

// 将两个有序数组合并成一个有序数组arr[left,mid] arr[mid+1,right]
func merge(arr []int, left, mid, right int) {

	//1.生成临时数组，定义左右指针
	temp := make([]int, right-left+1)
	i := left
	j := mid + 1
	k := 0

	//2.归并，两两比较将小的元素放入临时数组
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}
	//3.处理剩余的元素
	for i <= mid { //左边数组有剩余
		temp[k] = arr[i]
		i++
	}
	for j <= right {
		temp[k] = arr[j]
		j++
	}
	//4.将临时数组copy到原来的数组中
	for p := 0; p < len(temp); p++ {
		arr[p+left] = temp[p]
	}
}

// 将两个数组递归对半拆分
func mergeSort(arr []int, left, right int) {

	//递归终止的条件,当这个数组元素是1，或者空的时候，停止递归
	if left >= right {
		return
	}
	//1.计算中间值
	mid := left + (right-left)/2
	//2.递归拆分左半边部分
	mergeSort(arr, left, mid)
	//3.右边部分
	mergeSort(arr, mid+1, right)
	//4.调用合并函数，将两个有序的数组合并成一个有序的数组
	merge(arr, left, mid, right)
}

func MergeSort(arr []int) {
	n := len(arr)
	if n <= 1 { //不用排序
		return
	}
	mergeSort(arr, 0, n-1)
}
