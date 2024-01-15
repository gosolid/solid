//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSCoder : objc.NSObject
*/

type NSCoder struct {
  *objc.NSObject

}

func (r *NSCoder) EncodeValueOfObjCType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCoder) EncodeDataObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCoder) DecodeDataObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCoder) DecodeValueOfObjCType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCoder) VersionForClassName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

