//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLTileRenderPipelineColorAttachmentDescriptor : objc.NSObject
*/

type MTLTileRenderPipelineColorAttachmentDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLTileRenderPipelineColorAttachmentDescriptor) PixelFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineColorAttachmentDescriptor) SetPixelFormat() error {
  return fmt.Errorf("unimplemented")
}

