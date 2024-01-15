//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSRecursiveLock : objc.NSObject
*/

type NSRecursiveLock struct {
  *objc.NSObject

}

func (r *NSRecursiveLock) TryLock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRecursiveLock) LockBeforeDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRecursiveLock) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRecursiveLock) SetName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

