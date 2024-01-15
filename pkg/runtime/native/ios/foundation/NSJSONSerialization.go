//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSJSONSerialization : objc.NSObject
*/

type NSJSONSerialization struct {
  *objc.NSObject

}

func (r *NSJSONSerialization) JSONObjectWithData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSJSONSerialization) WriteJSONObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSJSONSerialization) JSONObjectWithStream() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSJSONSerialization) IsValidJSONObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSJSONSerialization) DataWithJSONObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

