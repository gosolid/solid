//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSMagnificationGestureRecognizer : AppKit.NSGestureRecognizer
*/

type NSMagnificationGestureRecognizer struct {
  *NSGestureRecognizer

}

func (r *NSMagnificationGestureRecognizer) SetMagnification() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMagnificationGestureRecognizer) Magnification() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

