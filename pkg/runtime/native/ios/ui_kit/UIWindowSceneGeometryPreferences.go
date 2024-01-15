//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UIWindowSceneGeometryPreferences : objc.NSObject
*/

type UIWindowSceneGeometryPreferences struct {
  *objc.NSObject

}

func (r *UIWindowSceneGeometryPreferences) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneGeometryPreferences) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

