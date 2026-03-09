package main

import (
	"Sort/bubbleSort"
	"Sort/selectionSort"
	"fmt"
)

func main() {
	fmt.Println("1111111")
	//冒泡测试
	arr1 := []int{5, 2, 9, 3, 7, 6, 1, 8, 4}
	arr2 := []int{9, 2, 3, 4, 5, 6, 7, 8, 1}
	bubbleSort.BubbleSortOptimize(arr1)
	bubbleSort.BubbleSortOptimize(arr2)
	fmt.Println("冒泡排序基础版本排序结果:", arr1)
	fmt.Println("冒泡排序改进版本排序结果:", arr2)

	//选择排序测试
	arr3 := []int{5, 2, 9, 3, 7, 6, 1, 8, 4}
	selectionSort.SelectionSort(arr3)
	fmt.Println("选择排序结果:", arr3)

	//测试不稳定性
	arr4 := []selectionSort.Num{
		{3, 0},
		{4, 1},
		{3, 2},
		{2, 3},
		{1, 4},
	}
	selectionSort.SelectionSortNun(arr4)
	fmt.Println("选择排序不稳定性结果:", arr4)
}
