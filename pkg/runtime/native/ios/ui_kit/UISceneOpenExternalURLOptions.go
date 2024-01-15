//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UISceneOpenExternalURLOptions : objc.NSObject
*/

type UISceneOpenExternalURLOptions struct {
  *objc.NSObject

}

func (r *UISceneOpenExternalURLOptions) UniversalLinksOnly() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneOpenExternalURLOptions) SetUniversalLinksOnly() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneOpenExternalURLOptions) EventAttribution() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneOpenExternalURLOptions) SetEventAttribution() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

