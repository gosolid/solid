//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLBlitPassDescriptor : objc.NSObject
*/

type MTLBlitPassDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLBlitPassDescriptor) BlitPassDescriptor() (*MTLBlitPassDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBlitPassDescriptor) SampleBufferAttachments() (*MTLBlitPassSampleBufferAttachmentDescriptorArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

