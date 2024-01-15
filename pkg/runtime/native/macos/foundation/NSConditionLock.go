//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSConditionLock : objc.NSObject
*/

type NSConditionLock struct {
  *objc.NSObject
  *NSLocking
}

func (r *NSConditionLock) InitWithCondition() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConditionLock) TryLock() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSConditionLock) UnlockWithCondition() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSConditionLock) LockBeforeDate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSConditionLock) LockWhenCondition() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSConditionLock) TryLockWhenCondition() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSConditionLock) Condition() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSConditionLock) Name() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConditionLock) SetName() error {
  return fmt.Errorf("unimplemented")
}

