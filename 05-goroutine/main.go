package main

import (
  "fmt"
  "time"
)



func infiniteHelloWorld(suffix string) {
  for {
    fmt.Println("Hello world", suffix)
    time.Sleep(0.1e9) // 0.1 seconds
  }
}





func main() {

  // "go" kicks of an internal thread/process/coroutine

  go infiniteHelloWorld("loop1")
  go infiniteHelloWorld("222222222222222 loop2")

  time.Sleep(10e9) // 10 seconds

}

