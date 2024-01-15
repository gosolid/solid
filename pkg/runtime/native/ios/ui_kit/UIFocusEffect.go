//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIFocusEffect : objc.NSObject
*/

type UIFocusEffect struct {
  *objc.NSObject

}

func (r *UIFocusEffect) Effect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusEffect) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusEffect) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

