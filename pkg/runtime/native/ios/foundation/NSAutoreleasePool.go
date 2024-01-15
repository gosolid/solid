//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSAutoreleasePool : objc.NSObject
*/

type NSAutoreleasePool struct {
  *objc.NSObject

}

func (r *NSAutoreleasePool) AddObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAutoreleasePool) Drain() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

