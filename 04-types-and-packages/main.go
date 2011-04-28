package main

import (
  "fmt"
  "somepackage"
  "somepackage2"
)


func main() {

  fmt.Println("main()")

  fmt.Println("--- testing somepackage ---")
  somepackage.SomeExternalFunc()
  // somepackage.someInternalFunc() // error: invalid reference to unexported identifier 'somepackage.someInternalFunc'


  fmt.Println("--- testing somepackage2 ---")
  testUser := somepackage2.CreateNormalUser("frode", "lol")
  fmt.Println(testUser)
  testUser2 := somepackage2.CreateAdminUser("frode2", "lol2")
  fmt.Println(testUser2)

  fmt.Println("testing with wrong password:", somepackage2.ValidateLogin(testUser2, "wrongpassword"))
  fmt.Println("testing with right password:", somepackage2.ValidateLogin(testUser2, "lo"+"l2"))

  var testUser2AsInterface somepackage2.S2UserInterface = testUser2
  fmt.Println("testing calling via interface", testUser2AsInterface.ValidateLogin("lol2"))

  fmt.Printf("The type of testUser2AsInterface is: %T\n", testUser2AsInterface)

  // Type assertions: (only works on interfaces, apparently).

  // thing, ok := testUser2.(somepackage2.ThingWithPassword) // error: type assertion only valid for interface types (???)
  thing, ok := (interface{}(testUser2)).(somepackage2.ThingWithPassword) // same as above but converting to empty interface first
  // thing, ok := testUser2AsInterface.(somepackage2.ThingWithPassword)  // similar to above but using already-converted interface
  fmt.Printf("Type type of the variable is %T and its value is %v. OK is: %v\n", thing, thing, ok)


}


