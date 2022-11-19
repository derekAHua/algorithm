package T100_

// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/

func sortedArrayToBST(nums []int) (ret *TreeNode) {
	if len(nums) == 0 {
		return nil
	}

	ret = &TreeNode{
		Val:   nums[len(nums)/2],
		Left:  sortedArrayToBST(nums[:len(nums)/2]),
		Right: sortedArrayToBST(nums[len(nums)/2+1:]),
	}

	return
}
