//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSArray : objc.NSObject
*/

type NSArray struct {
  *objc.NSObject
  *NSCopying
  *NSMutableCopying
  *NSSecureCoding
  *NSFastEnumeration
}

func (r *NSArray) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSArray) Count() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSArray) ObjectAtIndex() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSArray) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSArray) InitWithObjects() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

