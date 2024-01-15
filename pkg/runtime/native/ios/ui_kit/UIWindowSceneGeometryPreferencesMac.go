//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIWindowSceneGeometryPreferencesMac : UIKit.UIWindowSceneGeometryPreferences
*/

type UIWindowSceneGeometryPreferencesMac struct {
  *UIWindowSceneGeometryPreferences

}

func (r *UIWindowSceneGeometryPreferencesMac) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneGeometryPreferencesMac) InitWithSystemFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneGeometryPreferencesMac) SystemFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneGeometryPreferencesMac) SetSystemFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

