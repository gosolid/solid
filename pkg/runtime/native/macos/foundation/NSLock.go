//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSLock : objc.NSObject
*/

type NSLock struct {
  *objc.NSObject
  *NSLocking
}

func (r *NSLock) Name() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLock) SetName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLock) TryLock() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLock) LockBeforeDate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

