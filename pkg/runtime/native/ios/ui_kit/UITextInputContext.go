//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITextInputContext : objc.NSObject
*/

type UITextInputContext struct {
  *objc.NSObject

}

func (r *UITextInputContext) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInputContext) IsHardwareKeyboardInputExpected() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInputContext) IsDictationInputExpected() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInputContext) SetDictationInputExpected() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInputContext) SetHardwareKeyboardInputExpected() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInputContext) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInputContext) Current() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInputContext) IsPencilInputExpected() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInputContext) SetPencilInputExpected() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

