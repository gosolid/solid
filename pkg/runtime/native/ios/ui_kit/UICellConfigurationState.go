//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICellConfigurationState : UIKit.UIViewConfigurationState
*/

type UICellConfigurationState struct {
  *UIViewConfigurationState

}

func (r *UICellConfigurationState) SetExpanded() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellConfigurationState) IsSwiped() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellConfigurationState) CellDragState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellConfigurationState) SetCellDropState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellConfigurationState) IsEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellConfigurationState) SetEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellConfigurationState) IsReordering() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellConfigurationState) SetReordering() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellConfigurationState) SetCellDragState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellConfigurationState) CellDropState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellConfigurationState) IsExpanded() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellConfigurationState) SetSwiped() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

