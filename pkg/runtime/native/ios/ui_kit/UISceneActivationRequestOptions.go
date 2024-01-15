//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UISceneActivationRequestOptions : objc.NSObject
*/

type UISceneActivationRequestOptions struct {
  *objc.NSObject

}

func (r *UISceneActivationRequestOptions) RequestingScene() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneActivationRequestOptions) SetRequestingScene() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneActivationRequestOptions) CollectionJoinBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneActivationRequestOptions) SetCollectionJoinBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

