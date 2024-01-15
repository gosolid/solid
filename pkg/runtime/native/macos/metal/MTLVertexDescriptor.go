//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLVertexDescriptor : objc.NSObject
*/

type MTLVertexDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLVertexDescriptor) VertexDescriptor() (*MTLVertexDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexDescriptor) Reset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLVertexDescriptor) Layouts() (*MTLVertexBufferLayoutDescriptorArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexDescriptor) Attributes() (*MTLVertexAttributeDescriptorArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

