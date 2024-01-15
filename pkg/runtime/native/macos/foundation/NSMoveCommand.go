//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSMoveCommand : Foundation.NSScriptCommand
*/

type NSMoveCommand struct {
  *NSScriptCommand

}

func (r *NSMoveCommand) SetReceiversSpecifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMoveCommand) KeySpecifier() (*NSScriptObjectSpecifier, error) {
  return nil, fmt.Errorf("unimplemented")
}

