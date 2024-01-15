//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITextItem : objc.NSObject
*/

type UITextItem struct {
  *objc.NSObject

}

func (r *UITextItem) ContentType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextItem) Range() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextItem) Link() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextItem) TextAttachment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextItem) TagIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextItem) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextItem) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

