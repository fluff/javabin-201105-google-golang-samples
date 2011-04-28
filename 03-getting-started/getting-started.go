package main

import (
  "os"
  "fmt"
  //"math"
  "cmath"
)


func main() {

  fmt.Println("--------------- section 1 ----------------")

  var s1 string = "this is a string"
  fmt.Println(s1)
  
  // type inferred
  var s2 = "this is another string"
  fmt.Println(s2)
  
  // shorthand 
  s3 := "this is yet another string"
  fmt.Println(s3)

  // working with strings
  fmt.Println(s3[5] == 'i')
  // assigns a new string to variable, doesn't change contents string 
  s3 = "that could be a string" 
  fmt.Println(s3[5] == 'i')
  
  //   s3[5] = 'a' // error: invalid left hand side of argument









  fmt.Println("--------------- section 2 ----------------")

  // working with other native types
  var a  int    = 4                // machine-size (32bit or 64bit)
  var b  int8   = -128             // byte  (signed)
  var bu uint8  = 255              // byte  (unsigned)
  var w  int16  = -32768           // word  (16bits) (signed) 
  var wu uint16 = 65535            // word  (16bits) (unsigned) 
  var dw int32  = -1000000         // dword (32bits) (signed) 
  var du uint32 =  1000000         // dword (32bits) (unsigned) 
  var qw int64  = -100000000000000 // qword (64bits) (signed) 
  var qu uint64 =  100000000000000 // qword (64bits) (unsigned) 
  var f3 float32 = 3.1415999999999          // 32bit float
  var f6 float64 = 3.1415999999999          // 64bit float
  var c3 complex64  = complex(1.0, 1.0)     // 2x32bit complex float
  var c6 complex128 = complex(0.0, 1.0)     // 2x64bit complex float

  fmt.Println(a, b, bu, w, wu, dw, du, qw, qu)
  fmt.Println(a + 10)
  fmt.Println(qu + 10)
  fmt.Println(f3, f6, c3, c6)
  fmt.Println(c3*c3)
  fmt.Println(cmath.Sqrt(-1))
  //    fmt.Println(a + b) // error: incompatible types in binary expression



  // arrays
  var arrA [5]int = [5]int{1,2,3,5,8}   // or  = [...]int{1,2,3,5,8], compiler counts   
  fmt.Println(arrA)
  fmt.Println(arrA[3])
  var arrB [6]int
  //  arrB = arrA    // error: incompatible types in assignment
  var sliceB []int

  sliceB = arrA[:]     // converts array to slice
  fmt.Println(sliceB) 
  sliceB = arrA[1:3]   // converts array to slice
  fmt.Println(sliceB) 
  arrA[1] = 9999       // array are values, but slices are references
  fmt.Println(sliceB) 

  for i := 0; i < len(arrA); i++ {
    fmt.Println("loop test ",i,arrA[i])
  }
  for i, v := range arrA {
    fmt.Println("loop test ",i,v)
  }

  // maps
  var mapA map[string]int
  // mapA["test"] = 54   // panic: runtime error: invalid memory address or nil pointer dereference

  mapA = make(map[string]int) // maps are references, and must be allocated first
  mapA["test"] = 54
  mapA["hello"] = 985
  fmt.Println(mapA)



  fmt.Println("--------------- section 3 ----------------")


  // returning from main yields exit code 0
  os.Exit(123)

}
