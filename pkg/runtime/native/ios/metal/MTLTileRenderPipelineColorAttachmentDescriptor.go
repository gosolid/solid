//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLTileRenderPipelineColorAttachmentDescriptor : objc.NSObject
*/

type MTLTileRenderPipelineColorAttachmentDescriptor struct {
  *objc.NSObject

}

func (r *MTLTileRenderPipelineColorAttachmentDescriptor) PixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineColorAttachmentDescriptor) SetPixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

