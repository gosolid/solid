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
interface AppKit.NSScrubberLayout : objc.NSObject
*/

type NSScrubberLayout struct {
  *objc.NSObject
  *foundation.NSCoding
}

func (r *NSScrubberLayout) VisibleRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSScrubberLayout) ScrubberContentSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSScrubberLayout) LayoutAttributesForItemAtIndex() (*NSScrubberLayoutAttributes, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubberLayout) LayoutAttributesForItemsInRect() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubberLayout) Scrubber() (*NSScrubber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubberLayout) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubberLayout) InvalidateLayout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubberLayout) AutomaticallyMirrorsInRightToLeftLayout() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScrubberLayout) LayoutAttributesClass() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubberLayout) ShouldInvalidateLayoutForHighlightChange() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScrubberLayout) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubberLayout) PrepareLayout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubberLayout) ShouldInvalidateLayoutForChangeFromVisibleRect() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScrubberLayout) ShouldInvalidateLayoutForSelectionChange() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

