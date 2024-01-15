//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSInflectionRule : objc.NSObject
*/

type NSInflectionRule struct {
  *objc.NSObject

}

func (r *NSInflectionRule) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInflectionRule) AutomaticRule() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

