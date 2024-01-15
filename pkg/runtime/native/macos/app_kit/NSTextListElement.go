//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTextListElement : AppKit.NSTextParagraph
*/

type NSTextListElement struct {
  *NSTextParagraph

}

func (r *NSTextListElement) InitWithAttributedString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextListElement) TextListElementWithContents() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextListElement) TextList() (*NSTextList, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextListElement) MarkerAttributes() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextListElement) ChildElements() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextListElement) ParentElement() (*NSTextListElement, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextListElement) InitWithParentElement() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextListElement) Contents() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextListElement) AttributedString() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextListElement) TextListElementWithChildElements() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

