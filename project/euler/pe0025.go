package euler

import "math/big"

// https://projecteuler.net/problem=25

// PE0025 初めて1000桁になるFibonacci数の順番を計算する
func PE0025(n int) int {

	fibonacciMap := map[int]*big.Int{} // n番目のFibonacci数

	var fibonacci func(int) *big.Int

	fibonacci = func(nth int) *big.Int {
		if nth == 0 || nth == 1 {
			return big.NewInt(int64(1))
		} else if _, ok := fibonacciMap[nth]; ok {
			return fibonacciMap[nth]
		}
		fibonacciMap[nth] = big.NewInt(int64(0)).Add(fibonacci(nth-1), fibonacci(nth-2))
		if _, ok := fibonacciMap[nth-3]; ok {
			delete(fibonacciMap, nth-3)
		}
		return fibonacciMap[nth]
	}

	m := 0
	for calculateDigitsLength(fibonacci(m)) < n {
		m++
	}
	return m + 1
}

func calculateDigitsLength(n *big.Int) int {
	zero := big.NewInt(int64(0))
	ten := big.NewInt(int64(10))

	if n.Cmp(zero) == 0 {
		return 1
	}

	n.Abs(n)
	l := 0
	for ; n.Cmp(zero) > 0; n.Div(n, ten) {
		l++
	}
	return l
}
