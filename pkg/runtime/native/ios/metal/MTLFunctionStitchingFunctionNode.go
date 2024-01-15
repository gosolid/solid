//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLFunctionStitchingFunctionNode : objc.NSObject
*/

type MTLFunctionStitchingFunctionNode struct {
  *objc.NSObject

}

func (r *MTLFunctionStitchingFunctionNode) SetControlDependencies() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingFunctionNode) InitWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingFunctionNode) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingFunctionNode) SetName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingFunctionNode) Arguments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingFunctionNode) SetArguments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingFunctionNode) ControlDependencies() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

