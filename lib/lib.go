package lib

import (
	"strings"
)

// 找到等号
func FindEuq(inp string) int {
	for i, t := range inp {
		if t == '=' {
			return i
		}
	}
	return -1
}

func findAddSub(inp []string, start int) int {
	for i, t := range inp[start+1:] {
		if t == "+" || t == "-" {
			return i
		}
	}
	return -1
}

// 把方程转换成数组,方程的数字与符号用空格隔开
func EuqToArr(inpStr string) [][]string {
	inp := strings.Split(inpStr, " ")
	var res [][]string

	//截取+，-
	var(
		res = make([][]string, 0, len(inp)/2)
		end,oldEnd int
	)
	for i := 0; i < len(inp); i++ {
		if inp[i] == "+" || inp[i] == "-" {
			end = findAddSub(inp, i)
			res = append(res, inp[i:end])
			oldEnd=
		}
	}

	return res
}
