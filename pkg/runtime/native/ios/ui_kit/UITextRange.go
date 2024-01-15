//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITextRange : objc.NSObject
*/

type UITextRange struct {
  *objc.NSObject

}

func (r *UITextRange) IsEmpty() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextRange) Start() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextRange) End() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

