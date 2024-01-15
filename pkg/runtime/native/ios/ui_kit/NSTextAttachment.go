//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSTextAttachment : objc.NSObject
*/

type NSTextAttachment struct {
  *objc.NSObject

}

func (r *NSTextAttachment) SetFileType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) SetBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) AllowsTextAttachmentView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) TextAttachmentViewProviderClassForFileType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) SetContents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) SetImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) SetFileWrapper() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) UsesTextAttachmentView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) Bounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) LineLayoutPadding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) SetLineLayoutPadding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) InitWithData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) RegisterTextAttachmentViewProviderClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) Contents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) FileType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) Image() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) FileWrapper() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachment) SetAllowsTextAttachmentView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

