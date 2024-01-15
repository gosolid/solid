//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSFileSecurity : objc.NSObject
*/

type NSFileSecurity struct {
  *objc.NSObject

}

func (r *NSFileSecurity) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

