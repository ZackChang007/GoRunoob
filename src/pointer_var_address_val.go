package main



/*
指针变量 * 和地址值 & 的区别：
指针变量保存的是一个地址值，会分配独立的内存来存储一个整型数字。
当变量前面有 * 标识时，才等同于 & 的用法，否则会直接输出一个整型数字。
*/

// func main() {
// 	// var a int = 4
// 	a := 4
// 	var ptr *int
// 	ptr = &a
// 	println("a的值为", a);    // 4
// 	println("*ptr为", *ptr);  // 4
// 	println("ptr为", ptr);    // 824633794744
//  }


import "fmt"

 func main(){
    var a int = 4
    var ptr int
    ptr = a 
    fmt.Println(ptr)//4
    a = 15
    fmt.Println(ptr)//4
    
    var b int = 5 
    var ptr1 *int
    ptr1 = &b 
    fmt.Println(*ptr1)//5
    b=15 
    fmt.Println(*ptr1)//15
}