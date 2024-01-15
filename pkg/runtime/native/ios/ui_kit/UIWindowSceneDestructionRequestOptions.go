//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIWindowSceneDestructionRequestOptions : UIKit.UISceneDestructionRequestOptions
*/

type UIWindowSceneDestructionRequestOptions struct {
  *UISceneDestructionRequestOptions

}

func (r *UIWindowSceneDestructionRequestOptions) WindowDismissalAnimation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneDestructionRequestOptions) SetWindowDismissalAnimation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

