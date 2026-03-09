package bubbleSort

// 基础版本
func BubbleSortBasic(arr []int) {
	//先计算数组的长度
	n := len(arr)
	//由于最后一个数据自己就有序了，所以不用进行排序，只进行n-1轮比较
	for i := 0; i < n-1; i++ {
		//由于后面的i个数据已经排好了，所以要-i
		for j := 0; j < n-1-i; j++ {
			//如果前一个数据大于后一个数据，就调换
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 改进版，加个开关，解决数组有序还进行多余判断的情况
func BubbleSortOptimize(arr []int) {
	//计算数组的长度
	n := len(arr)
	//定义一个开关
	flat := false
	for i := 0; i < n-1; i++ {
		//每轮开始前将开关还原
		flat = false
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				//有顺序错误，交换并且打开开关
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flat = true
			}
		}
		//这轮比较如果没有数据转换，证明本身已经有序
		if !flat {
			break
		}
	}
	//	排序结束

}
