//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPasteControlConfiguration : objc.NSObject
*/

type UIPasteControlConfiguration struct {
  *objc.NSObject

}

func (r *UIPasteControlConfiguration) CornerStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteControlConfiguration) SetCornerStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteControlConfiguration) BaseBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteControlConfiguration) DisplayMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteControlConfiguration) SetDisplayMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteControlConfiguration) CornerRadius() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteControlConfiguration) SetCornerRadius() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteControlConfiguration) BaseForegroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteControlConfiguration) SetBaseForegroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteControlConfiguration) SetBaseBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

