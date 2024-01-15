//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIWindowSceneActivationRequestOptions : UIKit.UISceneActivationRequestOptions
*/

type UIWindowSceneActivationRequestOptions struct {
  *UISceneActivationRequestOptions

}

func (r *UIWindowSceneActivationRequestOptions) PreferredPresentationStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneActivationRequestOptions) SetPreferredPresentationStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneActivationRequestOptions) Placement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneActivationRequestOptions) SetPlacement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

