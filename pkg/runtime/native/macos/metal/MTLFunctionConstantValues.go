//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLFunctionConstantValues : objc.NSObject
*/

type MTLFunctionConstantValues struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLFunctionConstantValues) SetConstantValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLFunctionConstantValues) SetConstantValues() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLFunctionConstantValues) Reset() error {
  return fmt.Errorf("unimplemented")
}

