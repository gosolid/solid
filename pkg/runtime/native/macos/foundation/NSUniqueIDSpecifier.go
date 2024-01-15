//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUniqueIDSpecifier : Foundation.NSScriptObjectSpecifier
*/

type NSUniqueIDSpecifier struct {
  *NSScriptObjectSpecifier

}

func (r *NSUniqueIDSpecifier) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUniqueIDSpecifier) InitWithContainerClassDescription() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUniqueIDSpecifier) UniqueID() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUniqueIDSpecifier) SetUniqueID() error {
  return fmt.Errorf("unimplemented")
}

