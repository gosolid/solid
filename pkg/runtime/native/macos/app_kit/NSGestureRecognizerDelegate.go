//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSGestureRecognizerDelegate
*/

type NSGestureRecognizerDelegate struct {

}

func (r *NSGestureRecognizerDelegate) GestureRecognizer() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizerDelegate) GestureRecognizerShouldBegin() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

