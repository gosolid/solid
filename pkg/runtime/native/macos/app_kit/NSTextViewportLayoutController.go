//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSTextViewportLayoutController : objc.NSObject
*/

type NSTextViewportLayoutController struct {
  *objc.NSObject

}

func (r *NSTextViewportLayoutController) LayoutViewport() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) RelocateViewportToTextLocation() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) ViewportBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) ViewportRange() (*NSTextRange, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) AdjustViewportByVerticalOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) TextLayoutManager() (*NSTextLayoutManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) InitWithTextLayoutManager() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

