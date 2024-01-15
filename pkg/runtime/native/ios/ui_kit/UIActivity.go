//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UIActivity : objc.NSObject
*/

type UIActivity struct {
  *objc.NSObject

}

func (r *UIActivity) CanPerformWithActivityItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivity) PerformActivity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivity) ActivityDidFinish() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivity) ActivityCategory() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivity) ActivityTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivity) ActivityImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivity) PrepareWithActivityItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivity) ActivityType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivity) ActivityViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

