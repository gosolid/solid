//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSTextElement : objc.NSObject
*/

type NSTextElement struct {
  *objc.NSObject

}

func (r *NSTextElement) InitWithTextContentManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextElement) TextContentManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextElement) SetTextContentManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextElement) ElementRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextElement) SetElementRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextElement) ChildElements() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextElement) ParentElement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextElement) IsRepresentedElement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

