//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UIMotionEffect : objc.NSObject
*/

type UIMotionEffect struct {
  *objc.NSObject

}

func (r *UIMotionEffect) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMotionEffect) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMotionEffect) KeyPathsAndRelativeValuesForViewerOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

