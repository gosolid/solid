//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLRenderPipelineReflection : objc.NSObject
*/

type MTLRenderPipelineReflection struct {
  *objc.NSObject

}

func (r *MTLRenderPipelineReflection) MeshBindings() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineReflection) VertexArguments() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineReflection) FragmentArguments() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineReflection) TileArguments() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineReflection) VertexBindings() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineReflection) FragmentBindings() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineReflection) TileBindings() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineReflection) ObjectBindings() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

