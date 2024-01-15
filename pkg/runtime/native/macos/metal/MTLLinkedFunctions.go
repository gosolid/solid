//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Metal.MTLLinkedFunctions : objc.NSObject
*/

type MTLLinkedFunctions struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLLinkedFunctions) LinkedFunctions() (*MTLLinkedFunctions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLLinkedFunctions) SetFunctions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLLinkedFunctions) SetGroups() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLLinkedFunctions) Groups() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLLinkedFunctions) PrivateFunctions() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLLinkedFunctions) SetPrivateFunctions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLLinkedFunctions) Functions() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLLinkedFunctions) BinaryFunctions() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLLinkedFunctions) SetBinaryFunctions() error {
  return fmt.Errorf("unimplemented")
}

