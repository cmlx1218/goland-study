package main

import "fmt"

/**
 * @Author: CMLX
 * @Date: 2020/8/3 15:17
 * @Title: golang数据类型
 * --- --- ---
 * @Desc: 数据类型用于声明函数和对象
			布尔型：var b bool = true
			数据类型：整型：int和浮点型float32、float64
			字符串类型：
			派生类型：指针类型（Pointer）、数组类型、结构化类型（struct）、Channel类型、函数类型、切片类型、接口类型、Map类
1、数字类型
uint8、uint16、uint32、uint64 无符号x位整数
int8、int16、int32、int64 有符号x位整数
    2、浮点型
float32、float64 位浮点型数
complex64、compelex128 32位、64位实数和虚数
    3、布尔型
只可以是true和false  var b bool = false
    4、派生类型
指针类型（Pointer）
数组类型
结构化类型（struct）
Channel类型
函数类型
切片类型
接口类型（interface）
Map类型
    5、其他数字类型
byte：类型uint8
rune：类型uint32
uint：32或64位
int：与uint一样大小
uintptr：无符号整型，用于存放一个指针
*/

func main() {
	var a string

	a = "yyy"

	//a := "yyyy"	//报错：no new variables on left side of :=

	fmt.Println(a)

}
