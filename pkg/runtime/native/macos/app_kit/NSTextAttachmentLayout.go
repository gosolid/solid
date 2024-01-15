//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol AppKit.NSTextAttachmentLayout
*/

type NSTextAttachmentLayout struct {

}

func (r *NSTextAttachmentLayout) ImageForBounds() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentLayout) AttachmentBoundsForAttributes() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentLayout) ViewProviderForParentView() (*NSTextAttachmentViewProvider, error) {
  return nil, fmt.Errorf("unimplemented")
}

