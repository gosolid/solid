//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UISceneSession : objc.NSObject
*/

type UISceneSession struct {
  *objc.NSObject

}

func (r *UISceneSession) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSession) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSession) PersistentIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSession) StateRestorationActivity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSession) UserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSession) Scene() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSession) Role() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSession) Configuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSession) SetStateRestorationActivity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSession) SetUserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

