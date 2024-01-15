//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSAttributedStringMarkdownSourcePosition : objc.NSObject
*/

type NSAttributedStringMarkdownSourcePosition struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSAttributedStringMarkdownSourcePosition) InitWithStartLine() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownSourcePosition) RangeInString() (NSRange, error) {
  return NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownSourcePosition) StartLine() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownSourcePosition) StartColumn() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownSourcePosition) EndLine() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownSourcePosition) EndColumn() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

