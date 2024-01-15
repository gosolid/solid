//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSFormCell : AppKit.NSActionCell
*/

type NSFormCell struct {
  *NSActionCell

}

func (r *NSFormCell) IsOpaque() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFormCell) SetPreferredTextFieldWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFormCell) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFormCell) SetTitleFont() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFormCell) TitleWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFormCell) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFormCell) SetPlaceholderString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFormCell) SetTitleBaseWritingDirection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFormCell) PreferredTextFieldWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFormCell) InitTextCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFormCell) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFormCell) TitleFont() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFormCell) PlaceholderString() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFormCell) PlaceholderAttributedString() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFormCell) SetPlaceholderAttributedString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFormCell) TitleAlignment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFormCell) SetTitleAlignment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFormCell) InitImageCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFormCell) SetTitleWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFormCell) TitleBaseWritingDirection() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

