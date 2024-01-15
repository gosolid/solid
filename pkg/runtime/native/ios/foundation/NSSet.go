//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSSet : objc.NSObject
*/

type NSSet struct {
  *objc.NSObject

}

func (r *NSSet) Member() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSet) ObjectEnumerator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSet) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSet) InitWithObjects() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSet) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSet) Count() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

