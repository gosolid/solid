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
interface AppKit.NSTextContainer : objc.NSObject
*/

type NSTextContainer struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *NSTextLayoutOrientationProvider
}

func (r *NSTextContainer) Size() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) ExclusionPaths() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) MaximumNumberOfLines() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) SetWidthTracksTextView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) SetHeightTracksTextView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) LayoutManager() (*NSLayoutManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) TextLayoutManager() (*NSTextLayoutManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) SetMaximumNumberOfLines() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) WidthTracksTextView() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) InitWithSize() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) LineFragmentRectForProposedRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) SetSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) SetLineBreakMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) SetLineFragmentPadding() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) HeightTracksTextView() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) SetTextView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) ReplaceLayoutManager() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) SetExclusionPaths() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) LineBreakMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) LineFragmentPadding() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) IsSimpleRectangularTextContainer() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) TextView() (*NSTextView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContainer) SetLayoutManager() error {
  return fmt.Errorf("unimplemented")
}

