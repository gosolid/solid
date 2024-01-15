//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSURLCredential : objc.NSObject
*/

type NSURLCredential struct {
  *objc.NSObject
  *NSSecureCoding
  *NSCopying
}

func (r *NSURLCredential) Persistence() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

