//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UITextInputMode : objc.NSObject
*/

type UITextInputMode struct {
  *objc.NSObject

}

func (r *UITextInputMode) CurrentInputMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInputMode) PrimaryLanguage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInputMode) ActiveInputModes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

