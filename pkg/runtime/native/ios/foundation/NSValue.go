//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSValue : objc.NSObject
*/

type NSValue struct {
  *objc.NSObject

}

func (r *NSValue) GetValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSValue) InitWithBytes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSValue) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSValue) ObjCType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

