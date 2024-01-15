//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSOrderedSet : objc.NSObject
*/

type NSOrderedSet struct {
  *objc.NSObject

}

func (r *NSOrderedSet) InitWithObjects() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedSet) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedSet) Count() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedSet) ObjectAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedSet) IndexOfObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedSet) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

