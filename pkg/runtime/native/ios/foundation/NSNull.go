//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSNull : objc.NSObject
*/

type NSNull struct {
  *objc.NSObject

}

func (r *NSNull) Null() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

