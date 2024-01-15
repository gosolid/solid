//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSTextCheckingResult : objc.NSObject
*/

type NSTextCheckingResult struct {
  *objc.NSObject

}

func (r *NSTextCheckingResult) Range() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextCheckingResult) ResultType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

