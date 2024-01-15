//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSRangeSpecifier : Foundation.NSScriptObjectSpecifier
*/

type NSRangeSpecifier struct {
  *NSScriptObjectSpecifier

}

func (r *NSRangeSpecifier) StartSpecifier() (*NSScriptObjectSpecifier, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRangeSpecifier) SetStartSpecifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRangeSpecifier) EndSpecifier() (*NSScriptObjectSpecifier, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRangeSpecifier) SetEndSpecifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRangeSpecifier) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRangeSpecifier) InitWithContainerClassDescription() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

