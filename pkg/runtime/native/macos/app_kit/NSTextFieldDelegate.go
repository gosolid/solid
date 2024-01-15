//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSTextFieldDelegate
*/

type NSTextFieldDelegate struct {

}

func (r *NSTextFieldDelegate) TextField() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

