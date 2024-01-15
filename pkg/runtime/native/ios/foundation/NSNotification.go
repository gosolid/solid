//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSNotification : objc.NSObject
*/

type NSNotification struct {
  *objc.NSObject

}

func (r *NSNotification) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotification) Object() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotification) UserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotification) InitWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotification) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

