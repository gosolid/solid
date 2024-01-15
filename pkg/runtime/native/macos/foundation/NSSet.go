//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSSet : objc.NSObject
*/

type NSSet struct {
  *objc.NSObject
  *NSCopying
  *NSMutableCopying
  *NSSecureCoding
  *NSFastEnumeration
}

func (r *NSSet) Member() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSet) ObjectEnumerator() (*NSEnumerator, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSet) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSet) InitWithObjects() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSet) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSet) Count() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

