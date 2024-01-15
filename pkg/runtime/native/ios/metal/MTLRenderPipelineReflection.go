//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Metal.MTLRenderPipelineReflection : objc.NSObject
*/

type MTLRenderPipelineReflection struct {
  *objc.NSObject

}

func (r *MTLRenderPipelineReflection) ObjectBindings() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineReflection) MeshBindings() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineReflection) VertexArguments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineReflection) FragmentArguments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineReflection) TileArguments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineReflection) VertexBindings() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineReflection) FragmentBindings() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineReflection) TileBindings() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

