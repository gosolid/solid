//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIAlertController : UIKit.UIViewController
*/

type UIAlertController struct {
  *UIViewController

}

func (r *UIAlertController) Actions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertController) SetPreferredAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertController) TextFields() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertController) Message() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertController) AlertControllerWithTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertController) SetTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertController) PreferredStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertController) AddAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertController) AddTextFieldWithConfigurationHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertController) Severity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertController) SetSeverity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertController) PreferredAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertController) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAlertController) SetMessage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

