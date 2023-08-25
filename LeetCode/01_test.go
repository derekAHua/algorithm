package LeetCode

import (
	"math"
	"testing"
)

// @Author: Derek
// @Description:
// @Date: 2022/12/10 11:32
// @Version 1.0

func Test_weiBagProblem(t *testing.T) {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	println(weiBagProblem(weight, value, 4))
	println(weiBagProblem2(weight, value, 4))
}

var f = func(arr []int, l int, r int) int {
	pivot := arr[l]
	for l < r {
		for l < r && arr[r] >= pivot {
			r--
		}
		arr[l] = arr[r]
		for l < r && arr[l] <= pivot {
			l++
		}
		arr[r] = arr[l]
	}
	arr[l] = pivot
	return l
}

func quickSort(arr []int, l int, r int) {
	if l < r {
		p := f(arr, l, r)
		quickSort(arr, l, p-1)
		quickSort(arr, p+1, r)
	}
	return
}

func TestName(t *testing.T) {
	var theArray = []int{10, 9, 7, 5, 6, 3}
	quickSort(theArray, 0, len(theArray)-1)
	t.Log(theArray)
}

func TestA(t *testing.T) {
	t.Log(math.Cbrt(8))
}
