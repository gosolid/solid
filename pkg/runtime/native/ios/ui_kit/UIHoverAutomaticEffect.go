//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIHoverAutomaticEffect : objc.NSObject
*/

type UIHoverAutomaticEffect struct {
  *objc.NSObject

}

func (r *UIHoverAutomaticEffect) Effect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIHoverAutomaticEffect) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIHoverAutomaticEffect) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

