//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Metal.MTLVertexAttributeDescriptorArray : objc.NSObject
*/

type MTLVertexAttributeDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLVertexAttributeDescriptorArray) ObjectAtIndexedSubscript() (*MTLVertexAttributeDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttributeDescriptorArray) SetObject() error {
  return fmt.Errorf("unimplemented")
}

