package rel

import "math"

func Reliability(dataCount, parityCount uint8) string {
	var level string = "risky"

	// assume each node reliabilty is 0.9
	res := CalReliabilty(int(dataCount+parityCount), int(dataCount), 0.9)
	if res > 0.9999 {
		level = "High"
	} else if res > 0.99 {
		level = "Medium"
	}

	return level
}

// n is total: data+parity chunks
// m is data chunks
func CalReliabilty(n, m int, avail float64) float64 {
	res := float64(0)
	for i := 0; i < m; i++ {
		c := Binomial(uint64(n), uint64(i))
		a := math.Pow(avail, float64(i))
		una := float64(1) - avail
		una = math.Pow(una, float64(n-i))
		res += (float64(c) * a * una)
	}

	return 1 - res
}

func Binomial(n, k uint64) uint64 {
	if n < 0 || k < 0 {
		panic("negative input")
	}
	if n < k {
		panic("n must no less than k")
	}
	// (n,k) = (n, n-k)
	if k > n/2 {
		k = n - k
	}
	b := uint64(1)
	for i := uint64(1); i <= k; i++ {
		b = (n - k + i) * b / i
	}
	return b
}
