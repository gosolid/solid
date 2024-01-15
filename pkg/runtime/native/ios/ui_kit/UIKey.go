//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIKey : objc.NSObject
*/

type UIKey struct {
  *objc.NSObject

}

func (r *UIKey) Characters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKey) CharactersIgnoringModifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKey) ModifierFlags() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKey) KeyCode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

