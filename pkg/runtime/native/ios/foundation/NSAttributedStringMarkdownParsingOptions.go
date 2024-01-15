//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSAttributedStringMarkdownParsingOptions : objc.NSObject
*/

type NSAttributedStringMarkdownParsingOptions struct {
  *objc.NSObject

}

func (r *NSAttributedStringMarkdownParsingOptions) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) FailurePolicy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) SetLanguageCode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) AppliesSourcePositionAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) SetFailurePolicy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) LanguageCode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) SetAppliesSourcePositionAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) AllowsExtendedAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) SetAllowsExtendedAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) InterpretedSyntax() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) SetInterpretedSyntax() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

