//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLIndirectCommandBufferDescriptor : objc.NSObject
*/

type MTLIndirectCommandBufferDescriptor struct {
  *objc.NSObject

}

func (r *MTLIndirectCommandBufferDescriptor) InheritPipelineState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) MaxVertexBufferBindCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetMaxVertexBufferBindCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) MaxKernelThreadgroupMemoryBindCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetMaxKernelThreadgroupMemoryBindCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) MaxObjectBufferBindCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetCommandTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) MaxKernelBufferBindCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetMaxKernelBufferBindCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetSupportRayTracing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetSupportDynamicAttributeStride() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) CommandTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetInheritBuffers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetMaxFragmentBufferBindCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) MaxMeshBufferBindCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) MaxObjectThreadgroupMemoryBindCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetMaxObjectThreadgroupMemoryBindCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) InheritBuffers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) MaxFragmentBufferBindCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetMaxObjectBufferBindCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetMaxMeshBufferBindCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SupportRayTracing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SupportDynamicAttributeStride() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetInheritPipelineState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

