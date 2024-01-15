//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSPathControlDelegate
*/

type NSPathControlDelegate struct {

}

func (r *NSPathControlDelegate) PathControl() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

