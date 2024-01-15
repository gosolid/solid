//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTextElement : objc.NSObject
*/

type NSTextElement struct {
  *objc.NSObject

}

func (r *NSTextElement) SetTextContentManager() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextElement) ElementRange() (*NSTextRange, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextElement) SetElementRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextElement) ChildElements() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextElement) ParentElement() (*NSTextElement, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextElement) IsRepresentedElement() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextElement) InitWithTextContentManager() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextElement) TextContentManager() (*NSTextContentManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

