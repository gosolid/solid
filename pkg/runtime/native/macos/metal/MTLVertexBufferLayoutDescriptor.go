//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLVertexBufferLayoutDescriptor : objc.NSObject
*/

type MTLVertexBufferLayoutDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLVertexBufferLayoutDescriptor) StepRate() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLVertexBufferLayoutDescriptor) SetStepRate() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLVertexBufferLayoutDescriptor) Stride() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLVertexBufferLayoutDescriptor) SetStride() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLVertexBufferLayoutDescriptor) StepFunction() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLVertexBufferLayoutDescriptor) SetStepFunction() error {
  return fmt.Errorf("unimplemented")
}

