/*
Command line tool to try evaluating JSonnet.

Demo:
  echo "{ a: 1, b: 2 }"  | LD_LIBRARY_PATH=/usr/local/lib/ go run jsonnet_main/main.go  /dev/stdin
*/
package main

import "github.com/strickyak/jsonnet_cgo"

import (
  "fmt"
  "log"
  "os"
)

func main() {
  vm := jsonnet.Make()
  switch len(os.Args) {
  case 2:
    z, err := vm.EvaluateFile(os.Args[1])
    if err != nil {
      log.Fatalf("Error in jsonnet_main: %s", err)
    }
    fmt.Println(z)
  default:
    log.Fatal("Usage:  jsonnet_main filename")
  }
  vm.Destroy()
}
