//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSTokenFieldDelegate
*/

type NSTokenFieldDelegate struct {

}

func (r *NSTokenFieldDelegate) TokenField() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

