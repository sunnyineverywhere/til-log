# 문법 공부용 백준 문제 풀이 정리

## BOJ2420 사파리월드
```go
package main

import (
	"fmt"
)

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x-y
}


func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Print(absDiffInt(n, m))
}
```
두 개의 값을 입력받고 절대값을 계산하여 출력하는 문제.
기억할 포인트
- fmt.Scan
- 웬만하면 라이브러리를 가져다 쓰는 것보다, 함수를 직접 만드는게 빠르다(특히 Go는 체감상 더)
- 함수 형태에 익숙해지기
