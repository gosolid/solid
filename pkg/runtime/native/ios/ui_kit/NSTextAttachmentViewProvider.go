//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSTextAttachmentViewProvider : objc.NSObject
*/

type NSTextAttachmentViewProvider struct {
  *objc.NSObject

}

func (r *NSTextAttachmentViewProvider) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) TextLayoutManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) Location() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) InitWithTextAttachment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) LoadView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) AttachmentBoundsForAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) TextAttachment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) View() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) SetView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) TracksTextAttachmentViewBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) SetTracksTextAttachmentViewBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentViewProvider) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

