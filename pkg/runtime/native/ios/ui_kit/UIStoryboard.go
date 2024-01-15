//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIStoryboard : objc.NSObject
*/

type UIStoryboard struct {
  *objc.NSObject

}

func (r *UIStoryboard) StoryboardWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIStoryboard) InstantiateInitialViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIStoryboard) InstantiateInitialViewControllerWithCreator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIStoryboard) InstantiateViewControllerWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

