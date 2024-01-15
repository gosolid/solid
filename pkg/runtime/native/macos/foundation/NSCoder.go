//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSCoder : objc.NSObject
*/

type NSCoder struct {
  *objc.NSObject

}

func (r *NSCoder) EncodeValueOfObjCType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCoder) EncodeDataObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCoder) DecodeDataObject() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCoder) DecodeValueOfObjCType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCoder) VersionForClassName() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

