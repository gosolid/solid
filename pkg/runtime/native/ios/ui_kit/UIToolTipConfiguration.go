//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIToolTipConfiguration : objc.NSObject
*/

type UIToolTipConfiguration struct {
  *objc.NSObject

}

func (r *UIToolTipConfiguration) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIToolTipConfiguration) ToolTip() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIToolTipConfiguration) SourceRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIToolTipConfiguration) ConfigurationWithToolTip() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIToolTipConfiguration) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

