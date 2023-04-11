// https://www.runoob.com/go/go-data-types.html

// // go 1.9版本对于数字类型，无需定义int及float32、float64，系统会自动识别。
// package main
// import "fmt"

// func main() {
//    var a = 1.5
//    var b =2
//    fmt.Println(a,b)
// }


// 字符串去除空格和换行符
package main  
  
import (  
    "fmt"  
    "strings"  
)  
  
func main() {  
	// 下面str变量声明是以下完整声明方式的简化版：
	// var str = "这里是 www\n.runoob\n.com"
    str := "这里是 www\n.runoob\n.com"  
    fmt.Println("-------- 原字符串 ----------")  
    fmt.Println(str)  
    // 去除空格  
    str = strings.Replace(str, " ", "", -1)  
    // 去除换行符  
    str = strings.Replace(str, "\n", "", -1)  
    fmt.Println("-------- 去除空格与换行后 ----------")  
    fmt.Println(str)  
}