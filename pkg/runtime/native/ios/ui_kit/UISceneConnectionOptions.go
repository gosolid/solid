//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UISceneConnectionOptions : objc.NSObject
*/

type UISceneConnectionOptions struct {
  *objc.NSObject

}

func (r *UISceneConnectionOptions) URLContexts() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneConnectionOptions) UserActivities() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneConnectionOptions) CloudKitShareMetadata() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneConnectionOptions) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneConnectionOptions) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneConnectionOptions) NotificationResponse() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneConnectionOptions) ShortcutItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneConnectionOptions) SourceApplication() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneConnectionOptions) HandoffUserActivityType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

