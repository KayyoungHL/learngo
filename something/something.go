package something

import "fmt"

// 소문자로 시작하는 함수는 private 함수라서 외부에서 사용할 수 없다.
func sayBye() {
	fmt.Println("Bye")
}

// 대문자로 시작하는 함수는 public 함수라서 외부에서 사용할 수 있다.
func SayHello() {
	fmt.Println("Hello from something package")
}
