//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSString : objc.NSObject
*/

type NSString struct {
  *objc.NSObject
  *NSCopying
  *NSMutableCopying
  *NSSecureCoding
}

func (r *NSString) CharacterAtIndex() (uint16, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSString) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSString) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSString) Length() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

