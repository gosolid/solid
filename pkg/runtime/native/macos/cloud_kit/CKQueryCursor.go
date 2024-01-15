//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKQueryCursor : objc.NSObject
*/

type CKQueryCursor struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *CKQueryCursor) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQueryCursor) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

