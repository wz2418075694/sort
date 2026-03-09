package selectionSort

// 选择排序升序
func SelectionSort(arr []int) {
	//计算数组的长度
	n := len(arr)
	//arr[i]是未排序区的第一个元素
	for i := 0; i < n-1; i++ {
		//假设未排序区的第一个元素是最小值,将i的索引赋值给它
		index := i
		for j := i + 1; j < n; j++ {
			//如果有比这个数还小的值，将索引赋值给它
			if arr[j] < arr[index] {
				index = j
			}
			//一轮循环结束，将找出的最值和未排序区的第一个元素(arr[i])互换
			arr[i], arr[index] = arr[index], arr[i]
		}
	}
}

type Num struct {
	Value int //值
	Index int //索引
}

// 选择排序升序,结构体版本
func SelectionSortNun(arr []Num) {
	//计算数组的长度
	n := len(arr)
	//arr[i]是未排序区的第一个元素
	for i := 0; i < n-1; i++ {
		//假设未排序区的第一个元素是最小值,将i的索引赋值给它
		index := i
		for j := i + 1; j < n; j++ {
			//如果有比这个数还小的值，将索引赋值给它
			if arr[j].Value < arr[index].Value {
				index = j
			}
			//一轮循环结束，将找出的最值和未排序区的第一个元素(arr[i])互换
			arr[i], arr[index] = arr[index], arr[i]
		}
	}
}
