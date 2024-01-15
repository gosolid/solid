//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITextItemMenuPreview : objc.NSObject
*/

type UITextItemMenuPreview struct {
  *objc.NSObject

}

func (r *UITextItemMenuPreview) DefaultPreview() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextItemMenuPreview) InitWithView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextItemMenuPreview) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextItemMenuPreview) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

