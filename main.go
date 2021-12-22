package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/KayyoungHL/learngo/something"
)

// func 함수명(arg type, arg type, ...) return_type {
//     return value
// }
func multiply(a float64, b float64) float64 {
	return math.Round(a * b)
}

// 입력 argument가 같은 type일 경우 아래처럼 한 번만 적어주면 된다.
func divide(a, b float64) float64 {
	return math.Round(a / b)
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

// Naked return은 return 자리에 변수명과 타입을 같이 적어주고 return 할 때 return만 해주면 된다.
// 물론 return 자리에 원래 하던 방식으로 return 해도 동작한다.
func exampleNakedReturn(name string) (upperCase string) {
	defer fmt.Println("I'm done. Here it is.") // defer은 함수가 리턴하고 난 이후 동작한다. defer를 여러줄 사용할 경우 역순으로 동작한다?
	defer fmt.Println("놀랍게도 역순이네 ㅋㅋ")
	defer fmt.Println("그런데, 서순이 어떻게 되는거지?")
	defer fmt.Println("잘 되네. 좋다!")
	defer fmt.Println("이거 두 줄도 되나?")
	fmt.Println("I am making a Upper name")
	upperCase = strings.ToUpper(name)
	return
}

func superAdd(mode bool, numbers ...int) int {
	fmt.Println(numbers)
	total := 0
	// variable expression
	// if x := y+2 ; x > 10 {} 이런 느낌
	if mode {
		fmt.Println("For문 python 같은 느낌")
		for index, number := range numbers { // range는 index 먼저, 그 다음에 하나씩 꺼내서 넣어준다. python의 enumerate 같은 느낌
			fmt.Println(index, number)
			total += number
		}
	} else {
		fmt.Println("For문 자바나 c++ 같은 느낌")
		for i := 0; i < len(numbers); i++ {
			fmt.Println(i, numbers[i])
			total -= numbers[i]
		}
	}
	return total
}

func main() {
	fmt.Println("hello world")
	fmt.Println(math.Pi)
	// package를 만드는 방법
	// 1. go mod init
	// 2. go get 패키지 주소
	something.SayHello()

	// const => 상수. 바꿀 수 없음.
	// const const_name string = "Kay"

	// var => 변수. 바꿀 수 있음.
	var var_name string = "Kay"
	fmt.Println(var_name)

	// go는 기본적으로 type을 작성해 줘야 하는데 := 를 사용하면 알아서 type을 찾아서 넣어주면서 변수(var)로 지정해준다.
	// 축약형 := 를 사용하려면 반드시 func 내부에서 사용해야 한다.
	// func 외부에서는 축약형이 아닌 'var 변수명 type = 값' 또는 'const 상수명 type = 값' 으로 작성해야 한다.
	// 축약형 := 앞에 declared var 을 넣으면 "no new variables on left side of :=" 에러가 발생한다.

	name := "Kay"   // string을 입력하니 name은 자동으로 var name string 으로 type이 정해진다.
	age := 30       // int를 입력하니 age는 자동으로 var age int 로 type이 정해진다.
	height := 178.9 // float를 입력하니 height는 자동으로 var height float64 로 type이 정해진다.
	weight := 82.

	name = "Kayyoung"

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(multiply(height, weight))
	fmt.Println(divide(height, weight))

	totalLength, upperName := lenAndUpper(name)
	// go는 선언한 변수를 사용하지 않으면 "변수명 declared but not used" 라는 에러가 발생한다.
	// 성가시지만 fmt.Println을 해주자.
	fmt.Println(totalLength, upperName)

	// 만약 변수 두 개를 return 하는 함수를 썻을 때 하나만 쓰고 싶다면 _ 변수를 사용하자.
	// 이 변수는 compiler가 무시하는 변수이기 때문에 사용하지 않아도 에러가 발생하지 않는다.
	totalLength1, _ := lenAndUpper(var_name)
	fmt.Println(totalLength1)

	repeatMe("abc", "def", "ghi", "jkl", "mno", "pqr")

	fmt.Println(exampleNakedReturn(name))
	total := superAdd(true, 4, 6, 3, 6, 5, 1, 2, 7, 8, 5)
	fmt.Println(total)
	total -= superAdd(false, 4, 6, 3, 6, 5, 1, 2, 7, 8, 5)
	fmt.Println(total / 2)
}
