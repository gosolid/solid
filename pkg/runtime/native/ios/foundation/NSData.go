//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSData : objc.NSObject
*/

type NSData struct {
  *objc.NSObject

}

func (r *NSData) Length() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSData) Bytes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

