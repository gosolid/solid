//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSImageDelegate
*/

type NSImageDelegate struct {

}

func (r *NSImageDelegate) ImageDidNotDraw() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageDelegate) Image() error {
  return fmt.Errorf("unimplemented")
}

