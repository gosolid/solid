//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLBinaryArchiveDescriptor : objc.NSObject
*/

type MTLBinaryArchiveDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLBinaryArchiveDescriptor) Url() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBinaryArchiveDescriptor) SetUrl() error {
  return fmt.Errorf("unimplemented")
}

