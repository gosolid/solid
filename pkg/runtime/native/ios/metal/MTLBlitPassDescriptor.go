//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLBlitPassDescriptor : objc.NSObject
*/

type MTLBlitPassDescriptor struct {
  *objc.NSObject

}

func (r *MTLBlitPassDescriptor) SampleBufferAttachments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBlitPassDescriptor) BlitPassDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

