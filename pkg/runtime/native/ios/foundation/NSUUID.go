//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSUUID : objc.NSObject
*/

type NSUUID struct {
  *objc.NSObject

}

func (r *NSUUID) Compare() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUUID) UUIDString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUUID) UUID() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUUID) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUUID) InitWithUUIDString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUUID) InitWithUUIDBytes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUUID) GetUUIDBytes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

