//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSConditionLock : objc.NSObject
*/

type NSConditionLock struct {
  *objc.NSObject

}

func (r *NSConditionLock) TryLock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConditionLock) TryLockWhenCondition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConditionLock) LockBeforeDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConditionLock) Condition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConditionLock) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConditionLock) SetName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConditionLock) InitWithCondition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConditionLock) LockWhenCondition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConditionLock) UnlockWithCondition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

