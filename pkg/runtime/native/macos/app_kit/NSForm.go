//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSForm : AppKit.NSMatrix
*/

type NSForm struct {
  *NSMatrix

}

func (r *NSForm) DrawCellAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSForm) AddEntry() (*NSFormCell, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSForm) InsertEntry() (*NSFormCell, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSForm) SetPreferredTextFieldWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSForm) PreferredTextFieldWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSForm) SetInterlineSpacing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSForm) SetBezeled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSForm) SetTextFont() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSForm) SetTitleBaseWritingDirection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSForm) SetEntryWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSForm) RemoveEntryAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSForm) SetFrameSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSForm) SetTextAlignment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSForm) SetTitleFont() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSForm) CellAtIndex() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSForm) IndexOfCellWithTag() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSForm) SelectTextAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSForm) SetTextBaseWritingDirection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSForm) IndexOfSelectedItem() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSForm) SetBordered() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSForm) SetTitleAlignment() error {
  return fmt.Errorf("unimplemented")
}

