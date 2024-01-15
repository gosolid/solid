//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIScreenshotService : objc.NSObject
*/

type UIScreenshotService struct {
  *objc.NSObject

}

func (r *UIScreenshotService) WindowScene() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreenshotService) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreenshotService) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreenshotService) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreenshotService) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

