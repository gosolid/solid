//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface AppKit.NSStoryboard : objc.NSObject
*/

type NSStoryboard struct {
  *objc.NSObject

}

func (r *NSStoryboard) StoryboardWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStoryboard) InstantiateInitialController() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStoryboard) InstantiateInitialControllerWithCreator() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStoryboard) InstantiateControllerWithIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStoryboard) MainStoryboard() (*NSStoryboard, error) {
  return nil, fmt.Errorf("unimplemented")
}

