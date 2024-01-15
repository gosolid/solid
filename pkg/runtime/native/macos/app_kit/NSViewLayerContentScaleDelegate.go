//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSViewLayerContentScaleDelegate
*/

type NSViewLayerContentScaleDelegate struct {

}

func (r *NSViewLayerContentScaleDelegate) Layer() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

