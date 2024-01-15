//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSScrubberLayoutAttributes : objc.NSObject
*/

type NSScrubberLayoutAttributes struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSScrubberLayoutAttributes) ItemIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrubberLayoutAttributes) SetItemIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubberLayoutAttributes) Frame() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSScrubberLayoutAttributes) SetFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubberLayoutAttributes) Alpha() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrubberLayoutAttributes) SetAlpha() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubberLayoutAttributes) LayoutAttributesForItemAtIndex() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

