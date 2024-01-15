//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSNameSpecifier : Foundation.NSScriptObjectSpecifier
*/

type NSNameSpecifier struct {
  *NSScriptObjectSpecifier

}

func (r *NSNameSpecifier) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNameSpecifier) InitWithContainerClassDescription() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNameSpecifier) Name() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNameSpecifier) SetName() error {
  return fmt.Errorf("unimplemented")
}

