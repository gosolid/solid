//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSTextAttachmentViewProvider : objc.NSObject
*/

type NSTextAttachmentViewProvider struct {
  *objc.NSObject

}

func (r *NSTextAttachmentViewProvider) InitWithTextAttachment() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) TextLayoutManager() (*NSTextLayoutManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) TracksTextAttachmentViewBounds() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) SetTracksTextAttachmentViewBounds() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) LoadView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) AttachmentBoundsForAttributes() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) TextAttachment() (*NSTextAttachment, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) Location() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) View() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) SetView() error {
  return fmt.Errorf("unimplemented")
}

