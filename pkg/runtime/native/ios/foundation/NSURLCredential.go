//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSURLCredential : objc.NSObject
*/

type NSURLCredential struct {
  *objc.NSObject

}

func (r *NSURLCredential) Persistence() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

