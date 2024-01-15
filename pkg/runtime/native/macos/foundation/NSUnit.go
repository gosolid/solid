//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSUnit : objc.NSObject
*/

type NSUnit struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSUnit) InitWithSymbol() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnit) Symbol() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnit) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnit) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

