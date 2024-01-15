//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLCommandBufferDescriptor : objc.NSObject
*/

type MTLCommandBufferDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLCommandBufferDescriptor) RetainedReferences() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBufferDescriptor) SetRetainedReferences() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCommandBufferDescriptor) ErrorOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBufferDescriptor) SetErrorOptions() error {
  return fmt.Errorf("unimplemented")
}

