//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSPathCellDelegate
*/

type NSPathCellDelegate struct {

}

func (r *NSPathCellDelegate) PathCell() error {
  return fmt.Errorf("unimplemented")
}

