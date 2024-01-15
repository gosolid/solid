//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLCommandBufferDescriptor : objc.NSObject
*/

type MTLCommandBufferDescriptor struct {
  *objc.NSObject

}

func (r *MTLCommandBufferDescriptor) RetainedReferences() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBufferDescriptor) SetRetainedReferences() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBufferDescriptor) ErrorOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBufferDescriptor) SetErrorOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

