//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLIndirectCommandBufferDescriptor : objc.NSObject
*/

type MTLIndirectCommandBufferDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLIndirectCommandBufferDescriptor) SupportRayTracing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) CommandTypes() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetCommandTypes() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetInheritPipelineState() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetMaxFragmentBufferBindCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetMaxKernelThreadgroupMemoryBindCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SupportDynamicAttributeStride() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) InheritPipelineState() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetMaxVertexBufferBindCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) MaxFragmentBufferBindCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) MaxKernelBufferBindCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) MaxKernelThreadgroupMemoryBindCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetInheritBuffers() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetMaxKernelBufferBindCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetSupportRayTracing() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) SetSupportDynamicAttributeStride() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) InheritBuffers() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBufferDescriptor) MaxVertexBufferBindCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

