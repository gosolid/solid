//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UICalendarViewDecoration : objc.NSObject
*/

type UICalendarViewDecoration struct {
  *objc.NSObject

}

func (r *UICalendarViewDecoration) InitWithCustomViewProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarViewDecoration) DecorationWithColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarViewDecoration) DecorationWithImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarViewDecoration) DecorationWithCustomViewProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarViewDecoration) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarViewDecoration) InitWithImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

