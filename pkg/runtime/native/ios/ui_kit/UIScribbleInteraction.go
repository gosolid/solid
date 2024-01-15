//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIScribbleInteraction : objc.NSObject
*/

type UIScribbleInteraction struct {
  *objc.NSObject

}

func (r *UIScribbleInteraction) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScribbleInteraction) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScribbleInteraction) InitWithDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScribbleInteraction) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScribbleInteraction) IsHandlingWriting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScribbleInteraction) IsPencilInputExpected() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

