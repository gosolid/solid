//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSClickGestureRecognizer : AppKit.NSGestureRecognizer
*/

type NSClickGestureRecognizer struct {
  *NSGestureRecognizer
  *foundation.NSCoding
}

func (r *NSClickGestureRecognizer) NumberOfClicksRequired() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSClickGestureRecognizer) SetNumberOfClicksRequired() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSClickGestureRecognizer) NumberOfTouchesRequired() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSClickGestureRecognizer) SetNumberOfTouchesRequired() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSClickGestureRecognizer) ButtonMask() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSClickGestureRecognizer) SetButtonMask() error {
  return fmt.Errorf("unimplemented")
}

