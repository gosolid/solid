//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSData : objc.NSObject
*/

type NSData struct {
  *objc.NSObject
  *NSCopying
  *NSMutableCopying
  *NSSecureCoding
}

func (r *NSData) Bytes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSData) Length() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

