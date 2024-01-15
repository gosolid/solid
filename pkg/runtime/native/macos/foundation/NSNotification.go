//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSNotification : objc.NSObject
*/

type NSNotification struct {
  *objc.NSObject
  *NSCopying
  *NSCoding
}

func (r *NSNotification) UserInfo() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotification) InitWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotification) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotification) Name() (**NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNotification) Object() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

