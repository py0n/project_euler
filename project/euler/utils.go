package euler

import (
	"math/big"
)

// BigNumberLength 桁数を計算
func BigNumberLength(n *big.Int) int {
	m := big.NewInt(0).Set(n)
	m.Abs(m)
	return len(m.String())
}

// Gcm 最大公約數を計算する
func Gcm(n, m int) int {
	if n < 1 || m < 1 {
		return 0
	}
	if n < m {
		n, m = m, n
	}
	for {
		r := n % m
		if r > 0 {
			n, m = m, r
		}
		if r == 0 {
			return m
		}
	}
}

// IsPalindrome 囘文？
func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}

// Lcm 最小公倍數を計算する
func Lcm(n, m int) int {
	g := Gcm(n, m)
	if g == 0 {
		return 0
	} else if g == 1 {
		return n * m
	}
	return n * m / g
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

// DayOfWeek 曜日を表す型
type DayOfWeek int

// 曜日
const (
	None      DayOfWeek = iota // 該當無し
	Monday                     // 月曜日
	Tuesday                    // 火曜日
	Wednesday                  // 水曜日
	Thursday                   // 木曜日
	Friday                     // 金曜日
	Saturday                   // 土曜日
	Sunday                     // 日曜日
)

// ZellerWeekday ツェラーの公式
func ZellerWeekday(y, m, d int) DayOfWeek {
	if y < 1582 || (y == 1582 && m < 10) || (y == 1582 && m == 10 && d < 15) {
		return None
	}
	y_, m_ := y, m
	if m < 3 {
		y_, m_ = y-1, m+12
	}
	C := int(y_ / 100)
	G := 5*C + int(C/4)
	Y := y_ % 100
	return DayOfWeek(((d + int(26*(m_+1)/10) + Y + int(Y/4) + G + 5) % 7) + 1)
}

func multiplyDigits(n int) int {
	product := 1
	for n > 0 {
		product *= n % 10
		n = n / 10
	}
	return product
}
