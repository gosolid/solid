//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIWindowSceneProminentPlacement : UIKit.UIWindowScenePlacement
*/

type UIWindowSceneProminentPlacement struct {
  *UIWindowScenePlacement

}

func (r *UIWindowSceneProminentPlacement) ProminentPlacement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

