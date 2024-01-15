//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UIIndirectScribbleInteraction : objc.NSObject
*/

type UIIndirectScribbleInteraction struct {
  *objc.NSObject

}

func (r *UIIndirectScribbleInteraction) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIIndirectScribbleInteraction) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIIndirectScribbleInteraction) InitWithDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIIndirectScribbleInteraction) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIIndirectScribbleInteraction) IsHandlingWriting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

