//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSRecursiveLock : objc.NSObject
*/

type NSRecursiveLock struct {
  *objc.NSObject
  *NSLocking
}

func (r *NSRecursiveLock) TryLock() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRecursiveLock) LockBeforeDate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRecursiveLock) Name() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRecursiveLock) SetName() error {
  return fmt.Errorf("unimplemented")
}

