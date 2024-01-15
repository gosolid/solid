//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Metal.MTLType : objc.NSObject
*/

type MTLType struct {
  *objc.NSObject

}

func (r *MTLType) DataType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

