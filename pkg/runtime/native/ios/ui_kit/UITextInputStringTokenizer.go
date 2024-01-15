//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITextInputStringTokenizer : objc.NSObject
*/

type UITextInputStringTokenizer struct {
  *objc.NSObject

}

func (r *UITextInputStringTokenizer) InitWithTextInput() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

