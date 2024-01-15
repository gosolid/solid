//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol AppKit.NSTextAttachmentCell
*/

type NSTextAttachmentCell struct {

}

func (r *NSTextAttachmentCell) SetAttachment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentCell) DrawWithFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentCell) CellBaselineOffset() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentCell) TrackMouse() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentCell) CellSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentCell) WantsToTrackMouseForEvent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentCell) CellFrameForTextContainer() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentCell) Attachment() (*NSTextAttachment, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentCell) WantsToTrackMouse() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextAttachmentCell) Highlight() error {
  return fmt.Errorf("unimplemented")
}

