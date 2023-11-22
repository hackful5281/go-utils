package rand

import (
	"crypto/rand"
	"math/big"
	randM "math/rand"
	"time"
)

/**
math/rand          // 伪随机
crypto/rand        // 真随机
*/

// ============== 随机数 ==============

// RandNumberByMath 随机数
func RandNumberByMath() int {
	sec := randM.New(randM.NewSource(time.Now().UnixNano()))
	return sec.Int()
}

// RandRangeNumberByMath 带范围的随机数
func RandRangeNumberByMath(min int, max int) int {
	if min >= max {
		return -1
	}
	sec := randM.New(randM.NewSource(time.Now().UnixNano()))
	return sec.Intn(max-min) + min
}

// RandNumberByCrypto 随机数
func RandNumberByCrypto() int64 {
	safeNum, _ := rand.Int(rand.Reader, big.NewInt(time.Now().UnixNano()))
	return safeNum.Int64()
}

// RandRangeNumberByCrypto 带范围的随机数
func RandRangeNumberByCrypto(min int64, max int64) int64 {
	if min >= max || max <= 0 {
		return -1
	}
	safeNum, _ := rand.Int(rand.Reader, big.NewInt(max-min))
	return safeNum.Int64() + min
}

// ============== 随机字符串 ==============
var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomStringByMath 随机字符串
func RandomStringByMath(n int, allowedChars ...[]rune) string {
	sec := randM.New(randM.NewSource(time.Now().UnixNano()))
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[sec.Intn(len(letters))]
	}

	return string(b)
}

// RandomStringByCrypto 随机字符串
func RandomStringByCrypto(n int, allowedChars ...[]rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	size := len(letters)
	for i := range b {
		safeNum, _ := rand.Int(rand.Reader, big.NewInt(int64(size)))
		b[i] = letters[safeNum.Int64()]
	}

	return string(b)
}

// GenerateRandomNumber 生成 count 个[start, end] 结束的不重复的随机数
func GenerateRandomNumber(start int, end int, count int) []int {
	// 范围检查
	if end < start || (end-start) < count {
		return nil
	}

	// 存放结果的 slice
	nums := make([]int, 0)
	// 随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := randM.New(randM.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		// 生成随机数
		num := r.Intn(end-start) + start

		// 查看
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}
