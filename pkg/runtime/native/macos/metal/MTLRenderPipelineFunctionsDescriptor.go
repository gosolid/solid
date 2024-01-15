//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLRenderPipelineFunctionsDescriptor : objc.NSObject
*/

type MTLRenderPipelineFunctionsDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLRenderPipelineFunctionsDescriptor) VertexAdditionalBinaryFunctions() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineFunctionsDescriptor) SetVertexAdditionalBinaryFunctions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineFunctionsDescriptor) FragmentAdditionalBinaryFunctions() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineFunctionsDescriptor) SetFragmentAdditionalBinaryFunctions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineFunctionsDescriptor) TileAdditionalBinaryFunctions() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineFunctionsDescriptor) SetTileAdditionalBinaryFunctions() error {
  return fmt.Errorf("unimplemented")
}

