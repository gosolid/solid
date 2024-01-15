//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSInflectionRule : objc.NSObject
*/

type NSInflectionRule struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSInflectionRule) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInflectionRule) AutomaticRule() (*NSInflectionRule, error) {
  return nil, fmt.Errorf("unimplemented")
}

