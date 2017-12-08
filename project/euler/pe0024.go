package euler

// https://projecteuler.net/problem=23

// 先頭が `0` で固定されたときの順列の数は P(9,9) = 9! = 362,880 個。
// 先頭が `2` の数列は
//     1 * 362,880 * 2     =   725,761 番目から
//     1 * 362,880 * 3 - 1 = 1,088,639 番目まで。
//
// 先頭二文字が固定されたときの順列の数は P(8,8) = 8! = 40,320 個。
// `27` で始まる数列は
//     725,761 + 40,320 * 6     =   967,681 番目から
//     725,761 + 40,320 * 7 - 1 = 1,008,000 番目まで。
//
// 先頭三文字が固定されたときの順列の数は P(7,7) = 7! = 5,040 個。
// `278` で始まる数列は
//     967,681 + 5,040 * 6     =   997,921 番目から
//     967,681 + 5,040 * 7 - 1 = 1,002,960 番目まで。
//
// 先頭四文字が固定されたときの順列の数は P(6,6) = 6! = 720 個。
// `2783` で始まる数列は
//     997,921 + 720 * 2     =   999,361 番目から
//     997,921 + 720 * 3 - 1 = 1,000,080 番目まで。
//
// 先頭五文字が固定されたときの順列の数は P(5,5) = 5! = 120 個。
// `27839` で始まる数列は
//     999,361 + 120 * 5     =   999,961 番目から
//     999,361 + 120 * 6 - 1 = 1,000,080 番目まで。
//
// 先頭六文字が固定されたときの順列の数は P(4,4) = 4! = 24 個。
// `278391` で始まる数列は
//     999,961 + 24 * 1     =   999,985 番目から
//     999,961 + 24 * 2 - 1 = 1,000,008 番目まで。
//
// 先頭七文字が固定されたときの順列の数は P(3,3) = 3! = 6 個。
// `2783915` で始まる数列は
//     999,985 + 6 * 2     =   999,997 番目から
//     999,985 + 6 * 3 - 1 = 1,000,002 番目まで。
//
// 先頭八文字が固定されたときの順列の数は P(2,2) = 2! = 2 個。
// `27839154` で始まる数列は
//     999,997 + 2 * 1     =   999,999 番目から
//     999,997 + 2 * 2 - 1 = 1,000,000 番目まで。
//
// 結局 `2783915460` が 1,000,000 番目の数列。

// PE0024a 10個の数字の順列を辞書式に並べたときにn番目のものを計算する
func PE0024a(n int) string {
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	answer := []rune{}

	i := 1

	// n番目の順列を絞り込む
	for l := len(digits); l > 0; l = len(digits) {
		for k := 0; k < l; k++ {
			// 順列の数 P(n, n) = n!
			if p := factorial(l - 1); i+k*p <= n && n < i+(k+1)*p {
				answer = append(answer, digits[k])
				digits = append(digits[:k], digits[k+1:]...)
				i += k * p
				break
			}
		}
	}
	return string(answer)
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	product := n
	for i := n - 1; i > 0; i-- {
		product *= i
	}
	return product
}

// PE0024b digitsを入れ替えていくバージョン
func PE0024b(n int) string {
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	i, L := 1, len(digits)

	// n番目の順列を絞り込む
	for l := 0; l < L; l++ {
		for k := l; k < L; k++ {
			// 順列の数 P(n, n) = n!
			if p := factorial(L - l - 1); i+(k-l)*p <= n && n < i+(k-l+1)*p {
				for m := k; m > l; m-- {
					digits[m-1], digits[m] = digits[m], digits[m-1]
				}
				i += (k - l) * p
				break
			}
		}
	}
	return string(digits)
}
