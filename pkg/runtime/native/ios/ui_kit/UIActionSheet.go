//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIActionSheet : UIKit.UIView
*/

type UIActionSheet struct {
  *UIView

}

func (r *UIActionSheet) ShowFromToolbar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) ShowInView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) NumberOfButtons() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) CancelButtonIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) DestructiveButtonIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) SetDestructiveButtonIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) InitWithTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) ButtonTitleAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) ActionSheetStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) SetActionSheetStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) SetCancelButtonIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) FirstOtherButtonIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) ShowFromTabBar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) ShowFromBarButtonItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) ShowFromRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) IsVisible() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) AddButtonWithTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) DismissWithClickedButtonIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActionSheet) SetTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

