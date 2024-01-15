//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIPressesEvent : UIKit.UIEvent
*/

type UIPressesEvent struct {
  *UIEvent

}

func (r *UIPressesEvent) AllPresses() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPressesEvent) PressesForGestureRecognizer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

