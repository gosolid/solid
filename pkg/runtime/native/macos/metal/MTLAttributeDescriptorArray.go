//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Metal.MTLAttributeDescriptorArray : objc.NSObject
*/

type MTLAttributeDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLAttributeDescriptorArray) ObjectAtIndexedSubscript() (*MTLAttributeDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAttributeDescriptorArray) SetObject() error {
  return fmt.Errorf("unimplemented")
}

