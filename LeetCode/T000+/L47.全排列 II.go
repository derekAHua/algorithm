package T000_

// https://leetcode.cn/problems/permutations-ii/

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//示例 1：
//
//输入：nums = [1,1,2]
//输出： [[1,1,2], [1,2,1], [2,1,1]]
//示例 2：
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//提示：
//
//1 <= nums.length <= 8
//-10 <= nums[i] <= 10

func permuteUnique(nums []int) (ret [][]int) {
	var f func(map[int]bool, []int)

	f = func(m map[int]bool, arr []int) {
		if len(arr) > len(nums) {
			return
		}
		if len(arr) == len(nums) {
			r := make([]int, len(arr))
			copy(r, arr)
			ret = append(ret, r)
		}

		used := [21]bool{}
		for i := 0; i < len(nums); i++ {
			if m[i] || used[nums[i]+10] {
				continue
			}
			used[nums[i]+10] = true
			arr = append(arr, nums[i])
			m[i] = true
			f(m, arr)
			m[i] = false
			arr = arr[:len(arr)-1]
		}
	}

	f(map[int]bool{}, nil)
	return
}
