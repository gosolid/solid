//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSStackViewDelegate
*/

type NSStackViewDelegate struct {

}

func (r *NSStackViewDelegate) StackView() error {
  return fmt.Errorf("unimplemented")
}

