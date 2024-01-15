//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSPropertyListSerialization : objc.NSObject
*/

type NSPropertyListSerialization struct {
  *objc.NSObject

}

func (r *NSPropertyListSerialization) PropertyList() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyListSerialization) DataWithPropertyList() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyListSerialization) WritePropertyList() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyListSerialization) PropertyListWithData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyListSerialization) PropertyListWithStream() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyListSerialization) DataFromPropertyList() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyListSerialization) PropertyListFromData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

