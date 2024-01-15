//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIWindowSceneActivationAction : UIKit.UIAction
*/

type UIWindowSceneActivationAction struct {
  *UIAction

}

func (r *UIWindowSceneActivationAction) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneActivationAction) SetTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneActivationAction) ActionWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneActivationAction) ActionWithHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneActivationAction) ActionWithTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

