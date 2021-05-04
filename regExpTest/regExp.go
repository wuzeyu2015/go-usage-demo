package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abcazca7caac888a9ct ac"
	len := len(buf)
	//解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile(fmt.Sprintf("[0-9a-zA-Z_]{%d}", len))
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}

	match := reg1.Match([]byte(buf))
	fmt.Println("match = ", match)
	//根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result1 = ", result1)
}

