//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol AppKit.NSTextAttachmentContainer
*/

type NSTextAttachmentContainer struct {

}

func (r *NSTextAttachmentContainer) ImageForBounds() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentContainer) AttachmentBoundsForTextContainer() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

