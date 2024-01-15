//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKServerChangeToken : objc.NSObject
*/

type CKServerChangeToken struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *CKServerChangeToken) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKServerChangeToken) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

