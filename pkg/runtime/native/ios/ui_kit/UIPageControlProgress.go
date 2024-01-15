//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPageControlProgress : objc.NSObject
*/

type UIPageControlProgress struct {
  *objc.NSObject

}

func (r *UIPageControlProgress) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageControlProgress) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageControlProgress) CurrentProgress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageControlProgress) SetCurrentProgress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPageControlProgress) IsProgressVisible() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

