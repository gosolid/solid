//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Metal.MTLVertexBufferLayoutDescriptorArray : objc.NSObject
*/

type MTLVertexBufferLayoutDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLVertexBufferLayoutDescriptorArray) ObjectAtIndexedSubscript() (*MTLVertexBufferLayoutDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexBufferLayoutDescriptorArray) SetObject() error {
  return fmt.Errorf("unimplemented")
}

