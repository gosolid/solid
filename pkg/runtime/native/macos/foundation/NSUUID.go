//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSUUID : objc.NSObject
*/

type NSUUID struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSUUID) UUID() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUUID) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUUID) InitWithUUIDString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUUID) InitWithUUIDBytes() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUUID) GetUUIDBytes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUUID) Compare() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSUUID) UUIDString() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

