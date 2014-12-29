/*
jsonnet is a simple Go wrapper for the JSonnet VM.

See http://google.github.io/jsonnet/doc/index.html

Suggestions:
    go clone https://github.com/google/jsonnet.git
    cd jsonnet
    make libjsonnet.so
    cp libjsonnet.so /usr/local/lib/
    cp libjsonnet.h /usr/local/include/
*/
package jsonnet

/*
#include <stdio.h>
#include <libjsonnet.h>
#cgo LDFLAGS: -ljsonnet
*/
import "C"

import (
  "errors"
)

type VM struct {
  guts *C.struct_JsonnetVm
}

func Make() VM {
  return VM{guts: C.jsonnet_make()}
}

func (vm VM) Destroy() {
  C.jsonnet_destroy(vm.guts)
  vm.guts = nil
}

func (vm VM) EvaluateFile(filename string) (string, error) {
  var e C.int
  z := C.GoString(C.jsonnet_evaluate_file(vm.guts, C.CString(filename), &e))
  if e != 0 {
    return "", errors.New(z)
  }
  return z, nil
}
