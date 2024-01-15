//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Metal.MTLFunctionStitchingInputNode : objc.NSObject
*/

type MTLFunctionStitchingInputNode struct {
  *objc.NSObject
  *MTLFunctionStitchingNode
}

func (r *MTLFunctionStitchingInputNode) InitWithArgumentIndex() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingInputNode) ArgumentIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingInputNode) SetArgumentIndex() error {
  return fmt.Errorf("unimplemented")
}

