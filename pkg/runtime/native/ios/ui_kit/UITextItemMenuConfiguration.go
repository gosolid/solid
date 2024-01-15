//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITextItemMenuConfiguration : objc.NSObject
*/

type UITextItemMenuConfiguration struct {
  *objc.NSObject

}

func (r *UITextItemMenuConfiguration) ConfigurationWithMenu() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextItemMenuConfiguration) ConfigurationWithPreview() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextItemMenuConfiguration) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextItemMenuConfiguration) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

