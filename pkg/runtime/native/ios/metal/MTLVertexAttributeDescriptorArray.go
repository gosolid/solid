//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLVertexAttributeDescriptorArray : objc.NSObject
*/

type MTLVertexAttributeDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLVertexAttributeDescriptorArray) ObjectAtIndexedSubscript() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttributeDescriptorArray) SetObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

