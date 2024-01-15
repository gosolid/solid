//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSTextLayoutFragment : objc.NSObject
*/

type NSTextLayoutFragment struct {
  *objc.NSObject
  *foundation.NSSecureCoding
}

func (r *NSTextLayoutFragment) TextLineFragmentForVerticalOffset() (*NSTextLineFragment, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) FrameForTextAttachmentAtLocation() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) SetLayoutQueue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) State() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) LayoutFragmentFrame() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) TextAttachmentViewProviders() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) TextLayoutManager() (*NSTextLayoutManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) RangeInElement() (*NSTextRange, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) LayoutQueue() (*foundation.NSOperationQueue, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) BottomMargin() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) TextLineFragmentForTextLocation() (*NSTextLineFragment, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) InvalidateLayout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) RenderingSurfaceBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) TrailingPadding() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) TextElement() (*NSTextElement, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) TextLineFragments() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) LeadingPadding() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) TopMargin() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) InitWithTextElement() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) DrawAtPoint() error {
  return fmt.Errorf("unimplemented")
}

