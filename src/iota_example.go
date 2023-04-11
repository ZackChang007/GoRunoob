// package main

// import "fmt"

// func main() {
//     const (
//             a = iota   //0
//             b          //1
//             c          //2
//             d = "ha"   //独立值，iota += 1
//             e          //"ha"   iota += 1
//             f = 100    //iota +=1
//             g          //100  iota +=1
//             h = iota   //7,恢复计数
//             i          //8
//     )
//     fmt.Println(a,b,c,d,e,f,g,h,i)
// }


package main

import "fmt"
/*
iota 表示从 0 开始自动加 1，所以 i=1<<0, j=3<<1（<< 表示左移的意思），即：i=1, j=6，这没问题，关键在 k 和 l，从输出结果看 k=3<<2，l=3<<3。

简单表述:

i=1：左移 0 位，不变仍为 1。
j=3：左移 1 位，变为二进制 110，即 6。
注：python中转变bin为int的方法为：int('110',2)
k=3：左移 2 位，变为二进制 1100，即 12。
l=3：左移 3 位，变为二进制 11000，即 24。
注：<<n==*(2^n)。
*/
const (
    i=1<<iota
    j=3<<iota
    k
    l
)

func main() {
    fmt.Println("i=",i)
    fmt.Println("j=",j)
    fmt.Println("k=",k)
    fmt.Println("l=",l)
}