//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Metal.MTLFunctionStitchingGraph : objc.NSObject
*/

type MTLFunctionStitchingGraph struct {
  *objc.NSObject

}

func (r *MTLFunctionStitchingGraph) FunctionName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingGraph) Nodes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingGraph) Attributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingGraph) SetAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingGraph) InitWithFunctionName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingGraph) SetFunctionName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingGraph) SetNodes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingGraph) OutputNode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionStitchingGraph) SetOutputNode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

