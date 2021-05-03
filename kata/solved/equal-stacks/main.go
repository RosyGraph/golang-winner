package main

func equalStacks(h1, h2, h3 []int32) int32 {
	s1 := sum(h1)
	s2 := sum(h2)
	s3 := sum(h3)
	for s1 != s2 || s2 != s3 || s1 != s3 {
		if s1 > s2 && s1 > s3 {
			s1 -= h1[0]
			h1 = h1[1:]
		} else if s2 > s1 && s2 > s3 {
			s2 -= h2[0]
			h2 = h2[1:]
		} else if s3 > s1 && s3 > s2 {
			s3 -= h3[0]
			h3 = h3[1:]
		} else if s1 == s2 {
			if h1[0] > h2[0] {
				s1 -= h1[0]
				h1 = h1[1:]
			} else {
				s2 -= h2[0]
				h2 = h2[1:]
			}
		} else if s2 == s3 {
			if h2[0] > h3[0] {
				s2 -= h2[0]
				h2 = h2[1:]
			} else {
				s3 -= h3[0]
				h3 = h3[1:]
			}
		} else if s1 == s3 {
			if h1[0] > h3[0] {
				s1 -= h1[0]
				h1 = h1[1:]
			} else {
				s3 -= h3[0]
				h3 = h3[1:]
			}
		}
	}
	return sum(h1)
}

func sum(h []int32) int32 {
	var t int32
	for _, v := range h {
		t += v
	}
	return t
}
