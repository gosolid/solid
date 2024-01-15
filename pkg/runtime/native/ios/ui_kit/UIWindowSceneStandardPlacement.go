//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIWindowSceneStandardPlacement : UIKit.UIWindowScenePlacement
*/

type UIWindowSceneStandardPlacement struct {
  *UIWindowScenePlacement

}

func (r *UIWindowSceneStandardPlacement) StandardPlacement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

