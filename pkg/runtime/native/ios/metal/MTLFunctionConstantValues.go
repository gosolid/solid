//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLFunctionConstantValues : objc.NSObject
*/

type MTLFunctionConstantValues struct {
  *objc.NSObject

}

func (r *MTLFunctionConstantValues) SetConstantValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionConstantValues) SetConstantValues() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionConstantValues) Reset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

