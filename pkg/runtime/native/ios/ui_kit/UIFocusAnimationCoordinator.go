//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIFocusAnimationCoordinator : objc.NSObject
*/

type UIFocusAnimationCoordinator struct {
  *objc.NSObject

}

func (r *UIFocusAnimationCoordinator) AddCoordinatedUnfocusingAnimations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusAnimationCoordinator) AddCoordinatedAnimations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusAnimationCoordinator) AddCoordinatedFocusingAnimations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

