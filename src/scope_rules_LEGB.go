/*
https://www.runoob.com/go/go-scope-rules.html
可通过花括号来控制变量的作用域，花括号中的变量是单独的作用域，
同名变量会覆盖外层。
*/
package main

import "fmt"

var a = 5 // declare a using var

func main() {
    {
        a := 3 // use short variable declaration inside a function
        fmt.Println("in a = ", a)
    }
    fmt.Println("out a = ", a)
}