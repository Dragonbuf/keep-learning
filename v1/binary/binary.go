package binary

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
func singleNumberInThree(nums []int) int {

	var once, twice int
	for _, v := range nums {
		once = ^twice & (once ^ v)
		twice = ^once & (twice ^ v)
	}

	return once
}

//给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素
func singleNumber2(nums []int) []int {
	var mask int
	for _, v := range nums {
		mask ^= v
	}

	diff := mask & (-mask)

	x := 0
	for _, v := range nums {
		if v&diff != 0 {
			x ^= v
		}
	}

	return []int{x, mask ^ x}
}

//编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数（也被称为汉明重量）。
func hammingWeight(num uint32) int {
	var res int
	for num != 0 {
		num = num & (num - 1)
		res++
	}

	return res
}

////给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。

func countBits(num int) []int {
	var res []int
	for i := 0; i <= num; i++ {
		res = append(res, hammingWeight(uint32(i)))
	}
	return res
}

//颠倒给定的 32 位无符号整数的二进制位。
func reverseBits(num uint32) uint32 {
	res := uint32(0)
	pow := uint32(31)

	for num != 0 {
		res += (num & 1) << pow
		num >>= 1
		pow--
	}

	return res
}

func rangeBitwiseAnd(m int, n int) int {
	res := 0
	for m < n {
		m >>= 1
		n >>= 1
		res++
	}
	return m << res
}
