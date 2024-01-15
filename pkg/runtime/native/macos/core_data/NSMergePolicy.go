//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CoreData.NSMergePolicy : objc.NSObject
*/

type NSMergePolicy struct {
  *objc.NSObject

}

func (r *NSMergePolicy) RollbackMergePolicy() (*NSMergePolicy, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMergePolicy) MergeByPropertyObjectTrumpMergePolicy() (*NSMergePolicy, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMergePolicy) InitWithMergeType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMergePolicy) ResolveOptimisticLockingVersionConflicts() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMergePolicy) ResolveConstraintConflicts() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMergePolicy) OverwriteMergePolicy() (*NSMergePolicy, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMergePolicy) MergeByPropertyStoreTrumpMergePolicy() (*NSMergePolicy, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMergePolicy) MergeType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMergePolicy) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMergePolicy) ResolveConflicts() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMergePolicy) ErrorMergePolicy() (*NSMergePolicy, error) {
  return nil, fmt.Errorf("unimplemented")
}

