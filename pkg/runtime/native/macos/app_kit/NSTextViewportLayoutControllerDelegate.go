//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol AppKit.NSTextViewportLayoutControllerDelegate
*/

type NSTextViewportLayoutControllerDelegate struct {

}

func (r *NSTextViewportLayoutControllerDelegate) TextViewportLayoutControllerWillLayout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutControllerDelegate) TextViewportLayoutControllerDidLayout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutControllerDelegate) ViewportBoundsForTextViewportLayoutController() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTextViewportLayoutControllerDelegate) TextViewportLayoutController() error {
  return fmt.Errorf("unimplemented")
}

