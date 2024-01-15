//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLSharedTextureHandle : objc.NSObject
*/

type MTLSharedTextureHandle struct {
  *objc.NSObject

}

func (r *MTLSharedTextureHandle) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSharedTextureHandle) Label() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

