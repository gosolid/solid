//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSDeleteCommand : Foundation.NSScriptCommand
*/

type NSDeleteCommand struct {
  *NSScriptCommand

}

func (r *NSDeleteCommand) SetReceiversSpecifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDeleteCommand) KeySpecifier() (*NSScriptObjectSpecifier, error) {
  return nil, fmt.Errorf("unimplemented")
}

