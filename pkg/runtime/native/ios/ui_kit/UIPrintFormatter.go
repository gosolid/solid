//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPrintFormatter : objc.NSObject
*/

type UIPrintFormatter struct {
  *objc.NSObject

}

func (r *UIPrintFormatter) RequiresMainThread() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintFormatter) DrawInRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintFormatter) SetMaximumContentWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintFormatter) SetStartPage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintFormatter) SetPerPageContentInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintFormatter) PrintPageRenderer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintFormatter) MaximumContentHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintFormatter) SetContentInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintFormatter) RectForPageAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintFormatter) MaximumContentWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintFormatter) PerPageContentInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintFormatter) StartPage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintFormatter) PageCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintFormatter) RemoveFromPrintPageRenderer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintFormatter) SetMaximumContentHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPrintFormatter) ContentInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

