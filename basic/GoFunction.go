package main

/**
 * @Author: CMLX
 * @Date: 2020/8/4 14:30
 * @Title: GoLang函数
 * --- --- ---
 * @Desc:
 */

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	var result int = add(3, 4)
	fmt.Println(result)
}
