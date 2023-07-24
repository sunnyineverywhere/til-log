## 변수와 상수
### 1. 상수 선언
**일반적인 선언**
```go
const username = "lee"
const a int = 1
```

**괄호를 이용해 여러개의 값을 묶어서 초기화 가능**
```go
const (
  c1 = 10 // 첫 번째 값은 무조건 초기화
  c2
  c3 = "goorm" // 다른 변수형 선언 가능
)
```

### 2. 변수 선언
**키워드와 변수 형식 지정**
```go
var test bool = "ok
var name string = "sunny"
name2 := "sunny"
```
`:=` 는 `var ${변수명} string = ""`을 줄인 것

### 예제 - 간단한 덧셈
- var 키워드와 int 자료형을 모두 명시한 num1에 3을 선언 및 초기화합니다.
- := 연산자를 사용해 num2에 7을 선언 및 초기화합니다.
- := 연산자를 사용해 num1과 num2의 덧셈의 연산 및 결과 값을 result 변수에 선언 및 초기화 합니다.
- fmt 패키지의 Printf 함수를 한 번 사용해 결과값을 출력합니다. 
- (정수형 변수를 출력하는 서식 문자인 %d를 사용합니다.)

```go
package main

import "fmt"

func main() {
	var num1 int = 3
	num2 := 7
	result := num1 + num2
	fmt.Printf("%d과 %d의 합은 %d입니다.", num1, num2, result)
}
```

## 자료형 & 연산자
### 1. 연산자
+, -, *, /, % => 다른 건 특별히 없음

### 2. 자료형
- 불린 : bool
- 정수형: int
- 실수: float
- 복소수(신기하군,,,): complex
- 문자열: string

자료형 변환의 경우: `float(${다른 형의 변수})`와 같이 표기
강제 형변환도 가능

## 반복문
for문만 가능하다! 정도만 기억하면 될듯
```go
/* 일반적인 예 */
sum := 0
for i :=1; i<=10; i++ {
  sum += 1
}
```

```go
/* 확장된 for문의 예(신기하군,,,) */
n := 2
for n < 100 {
  fmt.Printf("count %d\n", n)
  n *= 2
}
```

```go
/* 확장된 for문의 예2 - 무한루프 */
n := 2
for {
		fmt.Printf("무한루프입니다.\n")
}
```
```go
/* for in range(파이썬...?과 유사) */
for index, num := range arr {
		fmt.Printf("arr[%d]의 값은 %d입니다.\n", index, num)
}
```

## 조건문
```go
	fmt.Print("정수입력 :")
	fmt.Scan(&num)

	if num == 1
	{
		fmt.Print("hello\n")
	}
	else if num == 2
	{
		fmt.Print("world\n")
	} else
	{
		fmt.Print("worng number..\n")
	}
```
거의 파이썬과 유사한데 `:` 표기만 없음

```go
if val := num * 2; val == 2 {
	fmt.Print("hello\n")
} else if val := num * 3; val == 6 {
	fmt.Print("world\n")
} else {
	fmt.Print("worng number..\n")
}
```
조건식 앞에(중괄호 앞에) 변수를 선언하고 식 입력 가능. 해당 조건문 블록에서만 사용할 수 있다(바로 뒤 중괄호까지가 범위인 듯).

스위치문은 그냥 자바랑 똑같아서 예시만 적어보고 넘어가겠음.
```go
switch(num){
	case 1:
		printf("1은 one\n");
		break;
	case 2:
		printf("2는 two\n");
		break;	
	case 3:
		printf("3은 three\n");
		break;
	default:
		printf("I don't know!");
}
```

후기
- 어차피... 백번 봐도 다시 까먹기 때문에^^... 분명히 공모전 하면서 다 봤던 건데 다시 보니 새롭다
- 어떤 건 자바같고 어떤 건 C같고 또 다시 보면 파이썬같다.
- 개인적으로 while문 없는 건 좋다. for를 확장해서 저렇게 쓴 건 누구 아이디어인지 참 대단하네
