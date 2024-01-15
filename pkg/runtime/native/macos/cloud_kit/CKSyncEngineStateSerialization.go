//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKSyncEngineStateSerialization : objc.NSObject
*/

type CKSyncEngineStateSerialization struct {
  *objc.NSObject
  *foundation.NSSecureCoding
}

func (r *CKSyncEngineStateSerialization) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineStateSerialization) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

