//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIWindowScenePlacement : objc.NSObject
*/

type UIWindowScenePlacement struct {
  *objc.NSObject

}

func (r *UIWindowScenePlacement) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowScenePlacement) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

