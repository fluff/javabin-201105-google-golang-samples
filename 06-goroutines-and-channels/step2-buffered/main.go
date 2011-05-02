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

  numberChannel := make(chan int, 100) // buffered channel
  go numberGenerator(3, numberChannel)

  for {
    time.Sleep(0.1e9)
    generatedNumber := <-numberChannel
    
  } 

}

