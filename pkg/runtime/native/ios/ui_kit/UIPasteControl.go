//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIPasteControl : UIKit.UIControl
*/

type UIPasteControl struct {
  *UIControl

}

func (r *UIPasteControl) InitWithConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteControl) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteControl) InitWithFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteControl) Configuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteControl) Target() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteControl) SetTarget() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

