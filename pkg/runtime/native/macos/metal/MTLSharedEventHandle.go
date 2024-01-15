//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Metal.MTLSharedEventHandle : objc.NSObject
*/

type MTLSharedEventHandle struct {
  *objc.NSObject
  *foundation.NSSecureCoding
}

func (r *MTLSharedEventHandle) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

