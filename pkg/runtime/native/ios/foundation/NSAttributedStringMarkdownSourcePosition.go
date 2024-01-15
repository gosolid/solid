//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSAttributedStringMarkdownSourcePosition : objc.NSObject
*/

type NSAttributedStringMarkdownSourcePosition struct {
  *objc.NSObject

}

func (r *NSAttributedStringMarkdownSourcePosition) InitWithStartLine() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownSourcePosition) RangeInString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownSourcePosition) StartLine() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownSourcePosition) StartColumn() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownSourcePosition) EndLine() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributedStringMarkdownSourcePosition) EndColumn() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

