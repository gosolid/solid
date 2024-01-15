//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLVertexAttributeDescriptor : objc.NSObject
*/

type MTLVertexAttributeDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLVertexAttributeDescriptor) SetBufferIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttributeDescriptor) Format() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttributeDescriptor) SetFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttributeDescriptor) Offset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttributeDescriptor) SetOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttributeDescriptor) BufferIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

