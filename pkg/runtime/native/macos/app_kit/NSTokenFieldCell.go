//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTokenFieldCell : AppKit.NSTextFieldCell
*/

type NSTokenFieldCell struct {
  *NSTextFieldCell

}

func (r *NSTokenFieldCell) SetCompletionDelay() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTokenFieldCell) DefaultCompletionDelay() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTokenFieldCell) TokenizingCharacterSet() (*foundation.NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTokenFieldCell) TokenStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTokenFieldCell) SetTokenStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTokenFieldCell) CompletionDelay() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTokenFieldCell) SetTokenizingCharacterSet() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTokenFieldCell) DefaultTokenizingCharacterSet() (*foundation.NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTokenFieldCell) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTokenFieldCell) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

