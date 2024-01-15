//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLFunctionStitchingInputNode : objc.NSObject
*/

type MTLFunctionStitchingInputNode struct {
  *objc.NSObject

}

func (r *MTLFunctionStitchingInputNode) InitWithArgumentIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingInputNode) ArgumentIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingInputNode) SetArgumentIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

