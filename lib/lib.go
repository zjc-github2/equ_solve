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

func find(inp []string, thing string, start int) int {
	for i, t := range inp[start+1:] {
		if t == thing {
			return i
		}
	}
	return -1
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

// 检查方程是否正确
func check(inp []string) {
	for _, t := range inp {
		if len(t) != 1 { //只有可能是数字
			t2 := []byte(t)
			for _, u := range t2 {
				if u < '0' || u > '9' { //有一个字符不是数字
					printErr("输入格式不对.")
				}
			}
		} else { //len(t)==1
			t2 := byte(t[0]) //反正只有一位
			//也不是加减乘除,也不是未知数
			isOper := (t2 == '+' || t2 == '-' || t2 == '*' || t2 == '/' || t2 == '^' || t2 == '(' || t2 == ')')
			isXYZ := ((t2 >= 'a' && t2 <= 'z') || (t2 >= 'A' && t2 <= 'Z'))
			if !isOper && !isXYZ {
				printErr("输入格式不对.")
			}
		}
	}
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

// 检查是不是数字
func isNum(inpStr string) bool {
	inp := []byte(inpStr)
	if inp[0] >= '0' && inp[0] <= '9' {
		return true
	}
	return false
}

// 两个括号相乘的情况
func NoKH_twoKH(inp1, inp2 []string) []string {
	var (
		pd1, pd2 = 1, 1 //判断是正还是负,后面一乘就行了
		res      []string
	)
	for i, t := range inp1 {
		if t == "+" {
			pd1 = 1
		} else if t == "-" {
			pd1 = -1
		}
		for j, u := range inp2 {
			if u == "+" {
				pd2 = 1
			} else if u == "-" {
				pd2 = -1
			}
			if pd1*pd2 == 1 {
				res = append(res, "+")
			} else { //pd1*pd2==-1
				res = append(res, "-")
			}

			if isNum(t) && isNum(u) {
				t2, err1 := strconv.Atoi(t)
				u2, err2 := strconv.Atoi(u)
				if err1 != nil || err2 != nil {
					panic("t2:" + err1.Error() + "\nu2:" + err2.Error())
				}
				res = append(res, strconv.Itoa(t2*u2*pd1*pd2))

				if inp1[i+1] == "*" && isXYZ(inp1[i+2]) {
					res = append(res, "*")
					res = append(res, inp1[i+2])
				}
				if inp2[j+1] == "*" && isXYZ(inp2[j+2]) {
					res = append(res, "*")
					res = append(res, inp2[j+2])
				}
			}
		}
	}

	return res
}

// 去括号
/*
func noKH(inp []string) []string {
	for i, t := range inp {
		if t[0] == '(' {
			end := find(inp, ")", i)
			sub := inp[i : end+1]

			if i == 0 || inp[i-1] == "+" { //i==0时相当于前面有个+
				if end == len(inp)-1 || inp[end+1] == "+" || inp[end+1] == "-" {
					inp[i] = ""
					inp[end] = ""
				} else if inp[end+1] == "*" {
					if inp[end+2] == "(" { //开始抽象
						end2 := find(inp, ")", end)
						sub2 := inp[end+2 : end2]
						for i, t := range sub { //分别相乘
							for i2, t2 := range sub2 {
								//TODO
							}
						}
					}
				}
			} else if inp[i-1] == "-" {
				inp[i] = ""
				inp[end] = ""
				for i, t := range sub {
					if isNum(t) {
						tNum, _ := strconv.Atoi(t)
						tNum = -tNum
						sub[i] = strconv.Itoa(tNum) //改sub会直接改到inp
					}
				}
			} else if inp[i-1] == "*" {

			}
		}
	}
}
*/
/*
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
*/
