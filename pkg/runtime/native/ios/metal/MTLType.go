//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLType : objc.NSObject
*/

type MTLType struct {
  *objc.NSObject

}

func (r *MTLType) DataType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

