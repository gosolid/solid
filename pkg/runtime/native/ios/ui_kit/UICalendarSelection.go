//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UICalendarSelection : objc.NSObject
*/

type UICalendarSelection struct {
  *objc.NSObject

}

func (r *UICalendarSelection) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarSelection) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarSelection) UpdateSelectableDates() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

