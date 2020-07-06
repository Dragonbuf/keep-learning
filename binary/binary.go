package binary

func rangeBitwiseAnd(m int, n int) int {
	res := 0
	for m < n {
		m >>= 1
		n >>= 1
		res++
	}
	return m << res
}
