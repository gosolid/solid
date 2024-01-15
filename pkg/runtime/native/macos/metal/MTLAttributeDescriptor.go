//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Metal.MTLAttributeDescriptor : objc.NSObject
*/

type MTLAttributeDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLAttributeDescriptor) Offset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAttributeDescriptor) SetOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAttributeDescriptor) BufferIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAttributeDescriptor) SetBufferIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAttributeDescriptor) Format() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAttributeDescriptor) SetFormat() error {
  return fmt.Errorf("unimplemented")
}

