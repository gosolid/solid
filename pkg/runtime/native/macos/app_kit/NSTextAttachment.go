//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSTextAttachment : objc.NSObject
*/

type NSTextAttachment struct {
  *objc.NSObject
  *NSTextAttachmentLayout
  *NSTextAttachmentContainer
  *foundation.NSSecureCoding
}

func (r *NSTextAttachment) SetFileType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) SetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) SetFileWrapper() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) SetLineLayoutPadding() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) AllowsTextAttachmentView() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) UsesTextAttachmentView() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) InitWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) RegisterTextAttachmentViewProviderClass() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) FileType() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) FileWrapper() (*foundation.NSFileWrapper, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) SetContents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) Bounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) SetBounds() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) AttachmentCell() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) InitWithFileWrapper() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) TextAttachmentViewProviderClassForFileType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) Contents() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) Image() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) SetAttachmentCell() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) LineLayoutPadding() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) SetAllowsTextAttachmentView() error {
  return fmt.Errorf("unimplemented")
}

