package main

import (
	"fmt"
)


func main() {
	var n int
    var paper [101][101]int
    var ans int = 0
	fmt.Scan(&n)

    for i := 0; i< n; i++ {
        var x, y int
        fmt.Scan(&x, &y)
        for j := 0; j < 10; j++ {
            for k := 0; k < 10; k++ {
                paper[x+j][y+k] = 1
            }
        }
    }

    for i:= 0; i< 101; i++ {
        for j := 0; j< 101; j++ {
            ans += paper[i][j]
        }
    }
	
    fmt.Print(ans)
}
