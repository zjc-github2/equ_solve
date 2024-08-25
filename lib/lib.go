package lib

import (
	"fmt"
	"os"
	"strconv"
)

func printErr(inp ...any) {
	fmt.Println(inp...)
	os.Exit(-1)
}

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

// 判断是不是未知数(字母)
func isXYZ(inpStr string) bool {
	inp := []byte(inpStr)
	if inp[0] >= 'A' && inp[0] <= 'Z' || inp[0] >= 'a' && inp[0] <= 'z' {
		if len(inp) != 1 {
			printErr("输入格式不对.")
		}
		return true
	}
	return false
}

// 检查是不是数字(主要检查输入有没有出错)
func isNum(inpStr string) bool {
	inp := []byte(inpStr)
	for _, t := range inp {
		if t < '0' || t > '9' {
			return false
		}
	}
	return true
}

// 把方程化为一般式
// 前提是没有括号，没有3*4这种可以化简的式子
// 比如3+2+2x=5+6(加法可以保留)
func GoToNormal(inp []string, equIndex int) []string {
	var (
		numAdd, xyzAdd int //各项系数,未知数系数之和
		plus           int //是正还是负,设为1/-1,后面一乘就行了
		//isFuHao        bool = true
	)
	//等号两边分别处理
	for i, t := range inp[:equIndex] {
		if len(t) == 1 && t[0] == '+' {
			plus = 1
			continue //不可能即时符号又是数字
		} else if len(t) == 1 && t[0] == '-' {
			plus = -1
			continue
		}
		if isNum(t) {
			num, err := strconv.Atoi(t)
			if err != nil {
				printErr("错误:", err)
			}
			numAdd += num * plus
		}
		//TODO:未知数
	}
}
