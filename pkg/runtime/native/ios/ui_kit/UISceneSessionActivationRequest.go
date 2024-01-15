//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UISceneSessionActivationRequest : objc.NSObject
*/

type UISceneSessionActivationRequest struct {
  *objc.NSObject

}

func (r *UISceneSessionActivationRequest) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSessionActivationRequest) Role() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSessionActivationRequest) Session() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSessionActivationRequest) Options() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSessionActivationRequest) SetOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSessionActivationRequest) Request() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSessionActivationRequest) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSessionActivationRequest) UserActivity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSessionActivationRequest) SetUserActivity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSessionActivationRequest) RequestWithRole() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSessionActivationRequest) RequestWithSession() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

