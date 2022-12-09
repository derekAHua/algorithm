package T100_

// https://programmercarl.com/0135.%E5%88%86%E5%8F%91%E7%B3%96%E6%9E%9C.html#%E6%80%9D%E8%B7%AF

func candy(ratings []int) (ret int) {
	c := make([]int, len(ratings))
	for i := range c {
		c[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] && c[i] <= c[i-1] {
			c[i] = c[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && c[i] <= c[i+1] {
			c[i] = c[i+1] + 1
		}
	}

	for _, v := range c {
		ret += v
	}
	return
}
