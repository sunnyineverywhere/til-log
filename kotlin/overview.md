

## 1강. 코틀린에서 변수를 다루는 방법



**변수 선언 키워드 - var과 val의 차이점**

| var | val |
|:---|:---|
| 가변 | 불변(==final) <br> List가 val로 선언되었을 때 element 자체는 추가 가능 |

<br>

타입을 의무적으로 작성하지 않아도 괜찮지만, 지정하고 싶은 경우에 다음과 같이 지정한다.
만약 변수를 최초 선언할 때, **변수에 초기값을 지정해주지 않는 경우에는 타입을 지정해야** 한다.

```kotlin
var number1: Long
```

<br>

변수를 선언할 때 primitive type과 reference type을 구분하지 않는다.
⇒ 프로그래머가 boxing과 unboxing을 구분하지 않아도 됨.

<br>

코틀린에서는 **‘null이 들어갈 수 있는’ 변수를 다르게 간주**한다.
```kotlin
var number1: Long? 
```

<br>

객체 인스턴스화를 할 때 new를 붙이지 않는다.
```kotlin
Person me = Person("name")
```

<br>


## 2강. 코틀린에서 null을 다루는 방법



```kotlin
func startsWithA1(str: String?) Boolean {
	if (str == null) {
		throw IllegalArgumentException("null이 들어왔다."
	}
	return str.startsWith("A")
}
```

<br>

**Safe call**

```kotlin
val str: String? = "ABC"

// 가능
str?.startsWith("A")

// 불가능
str.startsWith("A")
```

<br>

**Elvis 연산자**

```kotlin
val str: String? = null

str?.length ?: 0 // => 앞의 str?.length가 null이면 0 반환
```

<br>

**최종적으로 위의 코드 리팩토링**

```kotlin
func startsWithA1(str: String?) Boolean {
	return str?.startsWith("A") ?: throw IllegalArgumentException("null이 들어왔습니다")
}
```

<br>

**아무리 생각해도 null이 될 수 없는 경우**

```kotlin
fun startsWith(str: String?): Boolean {
		return str!!.startsWith("A")
}
```

<br>

## 3강. 코틀린에서 Type을 다루는 방법



**기본 타입**

`Kotlin` 암시적 타입 변경이 불가능함
↔ `Java` 암시적 타입 변경이 가능함

```kotlin
val number1 = 3
var number2 = number1.toDouble()

// 위와 같이 to{Type} 메서드를 사용해서 타입 변경을 한다
```

<br>

**타입 캐스팅**

```kotlin
fun print(obj: Any) {
	// is 의 반대 -> not is
	if (obj is Person) { // is == instanceof
	// as? -> 타입 캐스팅이 안될 경우 null 반환
		val person = obj as Person // -> 사실 필요하지 않음
		println(person.age)
	}
}
```

<br>

**Kotlin의 3가지 특이한 타입**

1) Any

- Java의 Object(모든 객체의 최상위 타입)
- 모든 Primitive Type의 최상의 타입도 Any
- Any 자체로는 null 포함 불가 → `Any?` 로 표현
- `equals()` `hashCode()` `toString()` 존재

2) Unit

- == Java ``void``
- 그 자체로 타입 인자로 사용 가능 (↔ ``void``)
- 함수형 프로그래밍에서 단 하나의 인스턴스만 갖는 타입을 의미
    
    ⇒ 코틀린에서, Unit은 실제 존재하는 타입이라는 것을 표현
    

3) Nothing

- 함수가 정상적으로 끝나지 않는 경우의 반환값

<br>

**String interpolation / String indexing**

```kotlin
val person = Person("이선의", 23)
println("이름 : ${person.name}")
```

⇒ 가독성, 일괄 변환, 정규식 활용

<br>

문자열 가공 시 사용

```kotlin
val str = """
ABC
EDFS
${person.name}
""".trimIndent()
```

<br>

문자열에서 문자를 가져올 때 index 사용 가능

```kotlin
val str = "ABC"
println(str[0])
```

<br>

## 4강. 코틀린에서 연산자를 다루는 방법



**단항 연산자 / 산술 연산자**

- 자바와 완전 동일
- 다만 객체를 비교할 때, **비교 연산자를 사용하면 자동으로 compareTo를 호출**

<br>

**비교 연산자와 동등성, 동일성**

- 동등성 `equals` ⇒ `==`
- 동일성  `==` ⇒ `===`

<br>

**논리 연산자**

- 자바와 완전히 동일 + Lazy 연산 수행
    - Lazy 연산
        
        ```kotlin
        if (fun2() || fun1()) {
         // fun2()만 true면 뒤의 연산을 수행하지 않음
        }
        ```
<br>     

**연산자 오버로딩**

```kotlin
val money1 = Money(1_000L)
val money2 = Money(2_000L)

print(money1 + money2) // -> 3000 출력
```
