//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSAttributedStringMarkdownParsingOptions : objc.NSObject
*/

type NSAttributedStringMarkdownParsingOptions struct {
  *objc.NSObject
  *NSCopying
}

func (r *NSAttributedStringMarkdownParsingOptions) SetFailurePolicy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) AppliesSourcePositionAttributes() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) SetAppliesSourcePositionAttributes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) AllowsExtendedAttributes() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) SetAllowsExtendedAttributes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) InterpretedSyntax() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) FailurePolicy() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) SetInterpretedSyntax() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) LanguageCode() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownParsingOptions) SetLanguageCode() error {
  return fmt.Errorf("unimplemented")
}

