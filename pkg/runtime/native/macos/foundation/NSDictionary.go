//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Foundation.NSDictionary : objc.NSObject
*/

type NSDictionary struct {
  *objc.NSObject
  *NSCopying
  *NSMutableCopying
  *NSSecureCoding
  *NSFastEnumeration
}

func (r *NSDictionary) InitWithObjects() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDictionary) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDictionary) Count() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDictionary) ObjectForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDictionary) KeyEnumerator() (*NSEnumerator, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDictionary) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

