//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLComputePipelineReflection : objc.NSObject
*/

type MTLComputePipelineReflection struct {
  *objc.NSObject

}

func (r *MTLComputePipelineReflection) Bindings() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineReflection) Arguments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

