//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIAlertView : UIKit.UIView
*/

type UIAlertView struct {
  *UIView

}

func (r *UIAlertView) IsVisible() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) AlertViewStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) InitWithFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) DismissWithClickedButtonIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) FirstOtherButtonIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) SetMessage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) SetAlertViewStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) InitWithTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) Show() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) SetTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) SetCancelButtonIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) NumberOfButtons() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) CancelButtonIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) AddButtonWithTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) ButtonTitleAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) TextFieldAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertView) Message() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

