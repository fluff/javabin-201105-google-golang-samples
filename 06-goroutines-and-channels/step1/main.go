package main

import (
  "fmt"
  "time"
)



func numberGenerator(multiplier int, outputChannel chan int) {
  for i := 0; ; i += multiplier {
    fmt.Println("Generating", i)
    outputChannel <- i       // blocks on writes to outputChannel
  }
}




func main() {

  numberChannel := make(chan int)
  go numberGenerator(3, numberChannel)

  time.Sleep(10e9) // 10 seconds

}

