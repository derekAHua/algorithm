package T100_

// https://leetcode.cn/problems/gas-station/

func canCompleteCircuit(gas []int, cost []int) (ret int) {
	sum := 0
	cur := 0
	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]

		cur += gas[i] - cost[i]
		if cur < 0 {
			ret = i + 1
			cur = 0
		}
	}

	if sum < 0 {
		return -1
	}
	return
}
