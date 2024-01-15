//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLComputePipelineReflection : objc.NSObject
*/

type MTLComputePipelineReflection struct {
  *objc.NSObject

}

func (r *MTLComputePipelineReflection) Bindings() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineReflection) Arguments() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

