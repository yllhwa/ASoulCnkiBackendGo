package check

import (
	"GoBackend/pkg/config"
	"math"
	"unicode/utf8"
)

func HashValue(s string, k int, q int64, b int) int64 {
	var num int64
	strRunes := []rune(s)
	length := utf8.RuneCountInString(s)
	for i := 0; i < length; i++ {
		num += int64(strRunes[i]) * int64(math.Pow(float64(b), float64(k-1-i)))
	}
	return num % q
}

func StringHash(s string, k int, q int64, b int, w int) []int64 {
	var _hash []int64
	strRunes := []rune(s)
	length := utf8.RuneCountInString(s)
	for i := 0; i < length-k+1; i++ {
		_hash = append(_hash, HashValue(string(strRunes[i:i+k]), k, q, b))
	}
	var hashPick []int64
	for i := 0; i+w <= len(_hash); i++ {
		var minHash int64
		for index, value := range _hash[i : i+w] {
			if index == 0 || value < minHash {
				minHash = value
			}
		}
		if len(hashPick) == 0 || minHash != hashPick[len(hashPick)-1] {
			hashPick = append(hashPick, minHash)
		}
	}
	return hashPick
}

func StringHashDefault(s string) []int64 {
	return StringHash(s, config.DefaultK, config.DefaultQ, config.DefaultB, config.DefaultW)
}
