//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Foundation.NSProtocolChecker : Foundation.NSProxy
*/

type NSProtocolChecker struct {
  *NSProxy

}

func (r *NSProtocolChecker) Protocol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProtocolChecker) Target() (*objc.NSObject, error) {
  return nil, fmt.Errorf("unimplemented")
}

