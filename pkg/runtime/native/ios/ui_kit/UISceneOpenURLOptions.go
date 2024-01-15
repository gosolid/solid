//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UISceneOpenURLOptions : objc.NSObject
*/

type UISceneOpenURLOptions struct {
  *objc.NSObject

}

func (r *UISceneOpenURLOptions) EventAttribution() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneOpenURLOptions) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneOpenURLOptions) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneOpenURLOptions) SourceApplication() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneOpenURLOptions) Annotation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneOpenURLOptions) OpenInPlace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

