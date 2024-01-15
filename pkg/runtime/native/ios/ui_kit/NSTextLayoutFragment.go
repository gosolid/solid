//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSTextLayoutFragment : objc.NSObject
*/

type NSTextLayoutFragment struct {
  *objc.NSObject

}

func (r *NSTextLayoutFragment) TextElement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) LeadingPadding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) TextLayoutManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) TextLineFragments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) LayoutFragmentFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) InitWithTextElement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) InvalidateLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) DrawAtPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) TopMargin() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) BottomMargin() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) TextLineFragmentForTextLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) RangeInElement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) LayoutQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) SetLayoutQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) State() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) RenderingSurfaceBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) TrailingPadding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) TextLineFragmentForVerticalOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) FrameForTextAttachmentAtLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLayoutFragment) TextAttachmentViewProviders() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

