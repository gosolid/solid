//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSActionCell : AppKit.NSCell
*/

type NSActionCell struct {
  *NSCell

}

func (r *NSActionCell) Tag() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSActionCell) SetTag() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSActionCell) Target() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSActionCell) SetTarget() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSActionCell) Action() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSActionCell) SetAction() error {
  return fmt.Errorf("unimplemented")
}

