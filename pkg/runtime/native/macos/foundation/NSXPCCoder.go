//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Foundation.NSXPCCoder : Foundation.NSCoder
*/

type NSXPCCoder struct {
  *NSCoder

}

func (r *NSXPCCoder) EncodeXPCObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCCoder) DecodeXPCObjectOfType() (**objc.NSObject, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCCoder) UserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCCoder) SetUserInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCCoder) Connection() (*NSXPCConnection, error) {
  return nil, fmt.Errorf("unimplemented")
}

