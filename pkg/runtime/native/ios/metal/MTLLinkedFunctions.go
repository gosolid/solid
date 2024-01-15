//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLLinkedFunctions : objc.NSObject
*/

type MTLLinkedFunctions struct {
  *objc.NSObject

}

func (r *MTLLinkedFunctions) BinaryFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLLinkedFunctions) SetBinaryFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLLinkedFunctions) SetGroups() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLLinkedFunctions) PrivateFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLLinkedFunctions) SetPrivateFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLLinkedFunctions) SetFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLLinkedFunctions) Functions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLLinkedFunctions) Groups() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLLinkedFunctions) LinkedFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

