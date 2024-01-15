//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSPropertyListSerialization : objc.NSObject
*/

type NSPropertyListSerialization struct {
  *objc.NSObject

}

func (r *NSPropertyListSerialization) WritePropertyList() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPropertyListSerialization) PropertyListWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyListSerialization) PropertyListWithStream() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyListSerialization) DataFromPropertyList() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyListSerialization) PropertyListFromData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyListSerialization) PropertyList() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPropertyListSerialization) DataWithPropertyList() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

