//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITextInputPasswordRules : objc.NSObject
*/

type UITextInputPasswordRules struct {
  *objc.NSObject

}

func (r *UITextInputPasswordRules) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInputPasswordRules) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInputPasswordRules) PasswordRulesWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInputPasswordRules) PasswordRulesDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

