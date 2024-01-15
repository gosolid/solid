//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Foundation.NSUnit : objc.NSObject
*/

type NSUnit struct {
  *objc.NSObject

}

func (r *NSUnit) InitWithSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnit) Symbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnit) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnit) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

