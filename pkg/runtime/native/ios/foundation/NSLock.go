//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSLock : objc.NSObject
*/

type NSLock struct {
  *objc.NSObject

}

func (r *NSLock) TryLock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLock) LockBeforeDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLock) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLock) SetName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

