//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSRelativeSpecifier : Foundation.NSScriptObjectSpecifier
*/

type NSRelativeSpecifier struct {
  *NSScriptObjectSpecifier

}

func (r *NSRelativeSpecifier) BaseSpecifier() (*NSScriptObjectSpecifier, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeSpecifier) SetBaseSpecifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRelativeSpecifier) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeSpecifier) InitWithContainerClassDescription() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeSpecifier) RelativePosition() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRelativeSpecifier) SetRelativePosition() error {
  return fmt.Errorf("unimplemented")
}

