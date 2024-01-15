//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSTextLayoutOrientationProvider
*/

type NSTextLayoutOrientationProvider struct {

}

func (r *NSTextLayoutOrientationProvider) LayoutOrientation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

