package main

func bestDivisor(n int32) int32 {
	var max int32
	var ans int32
	for i := int32(1); i < n+1; i++ {
		if n%i == 0 {
			digits := digitSum(i)
			if digits > max {
				max = digits
				ans = i
			} else if digits == max {
				if i < ans {
					ans = i
				}
			}
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
