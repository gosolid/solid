//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLSharedTextureHandle : objc.NSObject
*/

type MTLSharedTextureHandle struct {
  *objc.NSObject
  *foundation.NSSecureCoding
}

func (r *MTLSharedTextureHandle) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSharedTextureHandle) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

