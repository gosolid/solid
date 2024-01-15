//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSAutoreleasePool : objc.NSObject
*/

type NSAutoreleasePool struct {
  *objc.NSObject

}

func (r *NSAutoreleasePool) AddObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAutoreleasePool) Drain() error {
  return fmt.Errorf("unimplemented")
}

