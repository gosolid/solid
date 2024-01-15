//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLFunctionStitchingFunctionNode : objc.NSObject
*/

type MTLFunctionStitchingFunctionNode struct {
  *objc.NSObject
  *MTLFunctionStitchingNode
}

func (r *MTLFunctionStitchingFunctionNode) SetArguments() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingFunctionNode) ControlDependencies() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingFunctionNode) SetControlDependencies() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingFunctionNode) InitWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingFunctionNode) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingFunctionNode) SetName() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingFunctionNode) Arguments() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

