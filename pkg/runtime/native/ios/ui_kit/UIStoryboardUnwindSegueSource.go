//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIStoryboardUnwindSegueSource : objc.NSObject
*/

type UIStoryboardUnwindSegueSource struct {
  *objc.NSObject

}

func (r *UIStoryboardUnwindSegueSource) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIStoryboardUnwindSegueSource) SourceViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIStoryboardUnwindSegueSource) UnwindAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIStoryboardUnwindSegueSource) Sender() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

