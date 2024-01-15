//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIPickerView : UIKit.UIView
*/

type UIPickerView struct {
  *UIView

}

func (r *UIPickerView) NumberOfRowsInComponent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPickerView) SelectRow() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPickerView) SelectedRowInComponent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPickerView) ViewForRow() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPickerView) ReloadAllComponents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPickerView) NumberOfComponents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPickerView) RowSizeForComponent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPickerView) ReloadComponent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPickerView) SetShowsSelectionIndicator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPickerView) ShowsSelectionIndicator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPickerView) DataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPickerView) SetDataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPickerView) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPickerView) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

