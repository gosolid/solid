//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLAttributeDescriptor : objc.NSObject
*/

type MTLAttributeDescriptor struct {
  *objc.NSObject

}

func (r *MTLAttributeDescriptor) Format() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAttributeDescriptor) SetFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAttributeDescriptor) Offset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAttributeDescriptor) SetOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAttributeDescriptor) BufferIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAttributeDescriptor) SetBufferIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

