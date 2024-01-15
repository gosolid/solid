//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLAccelerationStructureDescriptor : objc.NSObject
*/

type MTLAccelerationStructureDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLAccelerationStructureDescriptor) Usage() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureDescriptor) SetUsage() error {
  return fmt.Errorf("unimplemented")
}

