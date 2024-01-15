//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSValue : objc.NSObject
*/

type NSValue struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSValue) ObjCType() (byte, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSValue) GetValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSValue) InitWithBytes() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSValue) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

