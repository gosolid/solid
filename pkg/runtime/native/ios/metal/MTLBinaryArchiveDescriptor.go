//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLBinaryArchiveDescriptor : objc.NSObject
*/

type MTLBinaryArchiveDescriptor struct {
  *objc.NSObject

}

func (r *MTLBinaryArchiveDescriptor) Url() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBinaryArchiveDescriptor) SetUrl() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

