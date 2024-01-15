//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.NSTextListElement : UIKit.NSTextParagraph
*/

type NSTextListElement struct {
  *NSTextParagraph

}

func (r *NSTextListElement) InitWithParentElement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextListElement) InitWithAttributedString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextListElement) TextListElementWithChildElements() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextListElement) Contents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextListElement) MarkerAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextListElement) TextListElementWithContents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextListElement) TextList() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextListElement) AttributedString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextListElement) ChildElements() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextListElement) ParentElement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

