//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSTextViewportLayoutController : objc.NSObject
*/

type NSTextViewportLayoutController struct {
  *objc.NSObject

}

func (r *NSTextViewportLayoutController) AdjustViewportByVerticalOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) TextLayoutManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) InitWithTextLayoutManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) ViewportBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) ViewportRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) LayoutViewport() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) RelocateViewportToTextLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutController) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

