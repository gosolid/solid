//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.NSTextParagraph : UIKit.NSTextElement
*/

type NSTextParagraph struct {
  *NSTextElement

}

func (r *NSTextParagraph) InitWithAttributedString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextParagraph) AttributedString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextParagraph) ParagraphContentRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextParagraph) ParagraphSeparatorRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

