//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLRenderPipelineFunctionsDescriptor : objc.NSObject
*/

type MTLRenderPipelineFunctionsDescriptor struct {
  *objc.NSObject

}

func (r *MTLRenderPipelineFunctionsDescriptor) SetVertexAdditionalBinaryFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineFunctionsDescriptor) FragmentAdditionalBinaryFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineFunctionsDescriptor) SetFragmentAdditionalBinaryFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineFunctionsDescriptor) TileAdditionalBinaryFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineFunctionsDescriptor) SetTileAdditionalBinaryFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineFunctionsDescriptor) VertexAdditionalBinaryFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

