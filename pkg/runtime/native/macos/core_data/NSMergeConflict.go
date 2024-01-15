//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CoreData.NSMergeConflict : objc.NSObject
*/

type NSMergeConflict struct {
  *objc.NSObject

}

func (r *NSMergeConflict) ObjectSnapshot() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMergeConflict) CachedSnapshot() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMergeConflict) PersistedSnapshot() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMergeConflict) NewVersionNumber() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMergeConflict) OldVersionNumber() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMergeConflict) InitWithSource() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMergeConflict) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMergeConflict) SourceObject() (*NSManagedObject, error) {
  return nil, fmt.Errorf("unimplemented")
}

