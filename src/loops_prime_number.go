// https://www.runoob.com/go/go-loops.html
// 求 100 内素数，我也来贴一个最终版的：
package main

import "fmt"

func main() {
    const RANGE = 100
    for num := 2; num <= RANGE; num++ {
        // 当前因数 factor 对应的另一个因数就是 num / factor。
        // 当前者大于后者时，可以认为所有因数已经分析完毕。
        for factor := 2; factor <= (num / factor); factor++ {
            //能除尽，则表示 factor 是 num 的一个因子。那么num就不是素数。
            if num%factor == 0 {
                goto NEXT_NUM
            }
        }
        fmt.Printf("%d\t是素数\n", num)
    NEXT_NUM:
    }
}