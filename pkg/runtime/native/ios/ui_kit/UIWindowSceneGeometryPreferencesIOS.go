//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIWindowSceneGeometryPreferencesIOS : UIKit.UIWindowSceneGeometryPreferences
*/

type UIWindowSceneGeometryPreferencesIOS struct {
  *UIWindowSceneGeometryPreferences

}

func (r *UIWindowSceneGeometryPreferencesIOS) InitWithInterfaceOrientations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneGeometryPreferencesIOS) InterfaceOrientations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneGeometryPreferencesIOS) SetInterfaceOrientations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneGeometryPreferencesIOS) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

