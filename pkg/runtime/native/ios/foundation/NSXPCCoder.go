//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSXPCCoder : Foundation.NSCoder
*/

type NSXPCCoder struct {
  *NSCoder

}

func (r *NSXPCCoder) SetUserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCCoder) Connection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCCoder) UserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

