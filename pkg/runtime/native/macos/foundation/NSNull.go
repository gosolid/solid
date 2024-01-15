//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSNull : objc.NSObject
*/

type NSNull struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSNull) Null() (*NSNull, error) {
  return nil, fmt.Errorf("unimplemented")
}

