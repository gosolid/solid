//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Metal.MTLFunctionStitchingGraph : objc.NSObject
*/

type MTLFunctionStitchingGraph struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLFunctionStitchingGraph) SetNodes() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingGraph) Attributes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingGraph) InitWithFunctionName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingGraph) SetFunctionName() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingGraph) Nodes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingGraph) SetAttributes() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingGraph) FunctionName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingGraph) OutputNode() (*MTLFunctionStitchingFunctionNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingGraph) SetOutputNode() error {
  return fmt.Errorf("unimplemented")
}

