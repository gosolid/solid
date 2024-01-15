//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLVertexAttributeDescriptor : objc.NSObject
*/

type MTLVertexAttributeDescriptor struct {
  *objc.NSObject

}

func (r *MTLVertexAttributeDescriptor) BufferIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttributeDescriptor) SetBufferIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttributeDescriptor) Format() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttributeDescriptor) SetFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttributeDescriptor) Offset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttributeDescriptor) SetOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

