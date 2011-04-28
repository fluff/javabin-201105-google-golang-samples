package somepackage2

import (
//  "fmt"
)

const AccessLevelNormal = 1
const AccessLevelAdmin  = 2

type S2User struct {
  username string
  password string
  accesslevel int
}

func CreateNormalUser(username, password string) *S2User {
  var newUser *S2User = new (S2User) // or: newUser := new(S2User)
  newUser.username = username
  newUser.password = password
  newUser.accesslevel = AccessLevelNormal
  return newUser
}

func CreateAdminUser(username, password string) *S2User {
  return &S2User{username, password, AccessLevelAdmin}
}

// Implementing a simple method taking the type as arg

func ValidateLogin(user *S2User, password string) bool {
  return (user.password == password)
}

// Implementing the equivalent method "on the type", note type between "func" and funcname
func (user *S2User) ValidateLogin(password string) bool {
  return ValidateLogin(user, password)
}


// Interfaces can be declared, any type that has matching methods automatically implement the interface
type S2UserInterface interface {
  ValidateLogin(password string) bool
}




type ThingWithPassword interface {
  RevealPassword() string
}

func (user *S2User) RevealPassword() string {
  return user.password
}



