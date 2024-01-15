//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSFileSecurity : objc.NSObject
*/

type NSFileSecurity struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSFileSecurity) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

