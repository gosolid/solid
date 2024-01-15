//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSJSONSerialization : objc.NSObject
*/

type NSJSONSerialization struct {
  *objc.NSObject

}

func (r *NSJSONSerialization) JSONObjectWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSJSONSerialization) WriteJSONObject() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSJSONSerialization) JSONObjectWithStream() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSJSONSerialization) IsValidJSONObject() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSJSONSerialization) DataWithJSONObject() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

