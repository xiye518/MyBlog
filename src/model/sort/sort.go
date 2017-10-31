package sort

import (
	"time"
	"math/rand"
)

func GetSlice() []int {
	var numbers []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 1000; i++ {
		a := r.Intn(100000)
		numbers = append(numbers, a)
	}
	return numbers
}

func 冒泡排序(array []int) ([]int) {
	for i := 0; i < len(array); i++ {
		for j := 1; j < len(array) - i; j++ {
			if array[j] < array[j - 1] {
				//交换
				array[j], array[j - 1] = array[j - 1], array[j]
			}
		}
	}
	return array
}

func 选择排序(array []int) ([]int) {
	length := len(array)
	for i := 0; i < length; i++ {
		maxIndex := 0
		//寻找最大的一个数，保存索引值
		for j := 1; j < length - i; j++ {
			if array[j] > array[maxIndex] {
				maxIndex = j
			}
		}
		array[length - i - 1], array[maxIndex] = array[maxIndex], array[length - i - 1]
	}
	return array
}



