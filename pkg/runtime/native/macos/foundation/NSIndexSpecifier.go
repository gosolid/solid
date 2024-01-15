//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSIndexSpecifier : Foundation.NSScriptObjectSpecifier
*/

type NSIndexSpecifier struct {
  *NSScriptObjectSpecifier

}

func (r *NSIndexSpecifier) InitWithContainerClassDescription() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexSpecifier) Index() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSIndexSpecifier) SetIndex() error {
  return fmt.Errorf("unimplemented")
}

