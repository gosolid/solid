//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSArray : objc.NSObject
*/

type NSArray struct {
  *objc.NSObject

}

func (r *NSArray) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSArray) InitWithObjects() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSArray) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSArray) Count() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSArray) ObjectAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

