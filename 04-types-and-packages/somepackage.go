package somepackage

import "fmt"

func SomeExternalFunc() {
  const loopCount = 3
  fmt.Println("Hello world from SomeExternalFunc")
  fmt.Println(someFuncWithArgs(loopCount, "hello loop "))
  result1, result2, result3 := someFuncWithMultiReturn(2)
  fmt.Println(result1)
  fmt.Println(result2)
  fmt.Println(result3)
}

func someInternalFunc() {
  fmt.Println("Hello world from someInternalFunc")
}

// arguments as usual, return type after arglist
func someFuncWithArgs(a int, b string) string {
  result := ""
  for i := 0; i<a; i++ {
    result += b
  }
  return result
}

func someFuncWithMultiReturn(a int) (string, string, int) {
  return "hello", "world", 42*a
}


