package main

func bestDivisor(n int32) int32 {
	var max int32
	var ans int32
	for i := int32(1); i < n+1; i++ {
		ans = 0
		if i < 10 {
			if i > max {
				max = i
				continue
			}
		}
		temp := i
		for temp > 9 {
			ans += temp % 10
			temp /= 10
		}
	}
	return ans
}

func digitSum(n int32) int32 {
	var res int32 = n % 10

	for n > 9 {
		n = n / 10
		res += n % 10
	}

	return res
}
