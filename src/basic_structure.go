// 当前程序的包名,一个可执行程序有且仅有一个 main 包。
package main

// 导入其他包,通过 import 关键字来导入其他非 main 包。
// import (
//     "fmt"
//     "math"
// )

// 前面加个点表示省略调用，那么调用该模块里面的函数Println，可以不用写模块名称了:
import . "fmt"

// // 为fmt起别名为fmt2
// import fmt2 "fmt"

// 常量定义
const PI = 3.14

// 全局变量的声明和赋值
var name = "gopher"

// 一般类型声明
type newType int

// 通过 type 关键字来进行结构(struct)和接口(interface)的声明。
// 结构的声明
type gopher struct{}

// 接口的声明
type golang interface{}

// Go语言中，使用大小写来决定该常量、变量、类型、接口、结构或函数是否可以被外部包所调用。
// 函数名首字母小写即为 private :
// func getId() {}
// 函数名首字母大写即为 public :
// func Printf() {}

// 由main函数作为程序入口点启动
func main() {
    Println("Hello World!")
}