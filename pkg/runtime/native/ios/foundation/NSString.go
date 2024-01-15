//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSString : objc.NSObject
*/

type NSString struct {
  *objc.NSObject

}

func (r *NSString) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSString) Length() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSString) CharacterAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSString) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

