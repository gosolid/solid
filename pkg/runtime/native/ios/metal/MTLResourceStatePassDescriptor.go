//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLResourceStatePassDescriptor : objc.NSObject
*/

type MTLResourceStatePassDescriptor struct {
  *objc.NSObject

}

func (r *MTLResourceStatePassDescriptor) ResourceStatePassDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLResourceStatePassDescriptor) SampleBufferAttachments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

