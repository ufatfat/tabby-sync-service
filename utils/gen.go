package utils

import "math/rand"

const toks = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ-_"

var tokBytes = []byte(toks)

func GenRandString(min, max int) string {
	if min > max {
		return ""
	}
	l := 0
	if min == max {
		l = min
	} else {
		l = rand.Intn(max-min) + min // 长度[10,25]
	}

	rst := ""
	for i := 0; i < l; i++ {
		rst += string(toks[rand.Intn(len(toks))])
	}
	return rst
}
