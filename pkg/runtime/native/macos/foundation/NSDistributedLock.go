//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSDistributedLock : objc.NSObject
*/

type NSDistributedLock struct {
  *objc.NSObject

}

func (r *NSDistributedLock) LockWithPath() (*NSDistributedLock, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDistributedLock) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDistributedLock) InitWithPath() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDistributedLock) TryLock() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDistributedLock) Unlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDistributedLock) BreakLock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDistributedLock) LockDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

