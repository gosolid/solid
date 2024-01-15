//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSSecureTextFieldCell : AppKit.NSTextFieldCell
*/

type NSSecureTextFieldCell struct {
  *NSTextFieldCell

}

func (r *NSSecureTextFieldCell) SetEchosBullets() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSecureTextFieldCell) EchosBullets() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

