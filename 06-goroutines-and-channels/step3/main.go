package main

import (
  "fmt"
  "time"
  "rand"
)



func numberGenerator(multiplier int, outputChannel chan int) {
  for i := 0; ; i += multiplier {
    fmt.Println(">>> Generating...", i)
    time.Sleep(rand.Int63n(10) * 1e8) // simulate CPU burn
    outputChannel <- i       // blocks on writes to outputChannel
  }
}

func numberProcessor(inputChannel chan int, outputChannel chan int) {
  for {
    generatedNumber := <-inputChannel
    fmt.Println("  <<< Processing...", generatedNumber)
    time.Sleep(rand.Int63n(10) * 1e8) // simulate CPU burn
    if (generatedNumber % 2 == 0) {
      outputChannel <- generatedNumber
    }
  }
}

func main() {

  numberChannel := make(chan int)
  go numberGenerator(3, numberChannel)
  resultChannel := make(chan int)
  go numberProcessor(numberChannel, resultChannel)

  for {
    fmt.Println("    === Received from processor:", <-resultChannel)
  } 

}

