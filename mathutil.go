package mathutility

// MaxInt ...
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MaxIntSlice ...
func MaxIntSlice(s []int) int {
	max := s[0]
	for _, n := range s {
		if n > max {
			max = n
		}
	}
	return max
}

// MinInt ...
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MinIntSlice ...
func MinIntSlice(s []int) int {
	min := s[0]
	for _, n := range s {
		if n < min {
			min = n
		}
	}
	return min
}
