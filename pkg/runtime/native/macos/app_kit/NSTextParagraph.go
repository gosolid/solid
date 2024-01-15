//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTextParagraph : AppKit.NSTextElement
*/

type NSTextParagraph struct {
  *NSTextElement

}

func (r *NSTextParagraph) InitWithAttributedString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextParagraph) AttributedString() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextParagraph) ParagraphContentRange() (*NSTextRange, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextParagraph) ParagraphSeparatorRange() (*NSTextRange, error) {
  return nil, fmt.Errorf("unimplemented")
}

