//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Metal.MTLBufferLayoutDescriptorArray : objc.NSObject
*/

type MTLBufferLayoutDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLBufferLayoutDescriptorArray) SetObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBufferLayoutDescriptorArray) ObjectAtIndexedSubscript() (*MTLBufferLayoutDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

