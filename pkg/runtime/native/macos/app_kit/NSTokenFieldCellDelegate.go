//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSTokenFieldCellDelegate
*/

type NSTokenFieldCellDelegate struct {

}

func (r *NSTokenFieldCellDelegate) TokenFieldCell() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

