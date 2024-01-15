//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLBufferLayoutDescriptorArray : objc.NSObject
*/

type MTLBufferLayoutDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLBufferLayoutDescriptorArray) ObjectAtIndexedSubscript() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBufferLayoutDescriptorArray) SetObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

