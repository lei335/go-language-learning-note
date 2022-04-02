package algo

// (n, k) combinatorial calculation; 0!=1
// 组合计算
func combical(n, k uint32) uint32 {
	if n < k {
		panic("parameter error, n shouldn't be less than k")
	}

	if k > (n - k) {
		k = n - k
	}

	b := uint32(1)
	for i := uint32(1); i <= k; i++ {
		b = (n - k + i) * b / i
	}
	return b
}
