//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLSharedEventListener : objc.NSObject
*/

type MTLSharedEventListener struct {
  *objc.NSObject

}

func (r *MTLSharedEventListener) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSharedEventListener) InitWithDispatchQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSharedEventListener) DispatchQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

