package euler

// IsPalindrome 囘文？
func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	r := []rune(s)
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-i-1] {
			return false
		}
	}
	return true
}

// Pow 累乘を計算
func Pow(m, n int) int {
	if n < 0 {
		return 0
	}
	p := 1
	for i := 0; i < n; i++ {
		p *= m
	}
	return p
}

// Pow10 10の累乘を計算
func Pow10(n int) int {
	return Pow(10, n)
}

// SumFactors 約數の総和を計算
func SumFactors(n int) int {
	if n < 2 {
		return 0
	}
	pm := PrimeFactorize(n)
	sum := 1
	for k, v := range pm {
		sum *= (Pow(k, (v+1)) - 1) / (k - 1)
	}
	return sum
}

func multiplyDigits(n int) int {
	product := 1
	for n > 0 {
		product *= n % 10
		n = n / 10
	}
	return product
}
