package main

import (
  "web"
)

func webCallback(arg string) string {
  return "hello web, your arg was "+arg
}

func main() {
  web.Get("/(.*)", webCallback)
  web.Run("0.0.0.0:9999")
}

