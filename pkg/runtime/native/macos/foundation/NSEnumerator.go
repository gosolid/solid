//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSEnumerator : objc.NSObject
*/

type NSEnumerator struct {
  *objc.NSObject
  *NSFastEnumeration
}

func (r *NSEnumerator) NextObject() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

