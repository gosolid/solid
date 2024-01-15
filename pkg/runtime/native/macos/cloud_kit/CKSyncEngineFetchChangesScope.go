//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface CloudKit.CKSyncEngineFetchChangesScope : objc.NSObject
*/

type CKSyncEngineFetchChangesScope struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *CKSyncEngineFetchChangesScope) InitWithZoneIDs() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchChangesScope) InitWithExcludedZoneIDs() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchChangesScope) ZoneIDs() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchChangesScope) ExcludedZoneIDs() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

