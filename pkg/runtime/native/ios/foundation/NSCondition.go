//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSCondition : objc.NSObject
*/

type NSCondition struct {
  *objc.NSObject

}

func (r *NSCondition) Wait() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCondition) WaitUntilDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCondition) Signal() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCondition) Broadcast() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCondition) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCondition) SetName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

