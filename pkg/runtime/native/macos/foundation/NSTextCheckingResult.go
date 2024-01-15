//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSTextCheckingResult : objc.NSObject
*/

type NSTextCheckingResult struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSTextCheckingResult) ResultType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingResult) Range() (NSRange, error) {
  return NSRange{}, fmt.Errorf("unimplemented")
}

