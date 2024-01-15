//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSOrderedSet : objc.NSObject
*/

type NSOrderedSet struct {
  *objc.NSObject
  *NSCopying
  *NSMutableCopying
  *NSSecureCoding
  *NSFastEnumeration
}

func (r *NSOrderedSet) IndexOfObject() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOrderedSet) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedSet) InitWithObjects() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedSet) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedSet) Count() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOrderedSet) ObjectAtIndex() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

