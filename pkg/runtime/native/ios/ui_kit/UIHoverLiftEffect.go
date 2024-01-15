//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIHoverLiftEffect : objc.NSObject
*/

type UIHoverLiftEffect struct {
  *objc.NSObject

}

func (r *UIHoverLiftEffect) Effect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIHoverLiftEffect) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIHoverLiftEffect) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

