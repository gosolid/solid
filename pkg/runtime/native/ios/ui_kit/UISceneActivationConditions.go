//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UISceneActivationConditions : objc.NSObject
*/

type UISceneActivationConditions struct {
  *objc.NSObject

}

func (r *UISceneActivationConditions) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneActivationConditions) CanActivateForTargetContentIdentifierPredicate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneActivationConditions) SetCanActivateForTargetContentIdentifierPredicate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneActivationConditions) PrefersToActivateForTargetContentIdentifierPredicate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneActivationConditions) SetPrefersToActivateForTargetContentIdentifierPredicate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneActivationConditions) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

