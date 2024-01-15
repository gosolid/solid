//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSCondition : objc.NSObject
*/

type NSCondition struct {
  *objc.NSObject
  *NSLocking
}

func (r *NSCondition) SetName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCondition) Wait() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCondition) WaitUntilDate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCondition) Signal() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCondition) Broadcast() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCondition) Name() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

