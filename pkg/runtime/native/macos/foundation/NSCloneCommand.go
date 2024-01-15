//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSCloneCommand : Foundation.NSScriptCommand
*/

type NSCloneCommand struct {
  *NSScriptCommand

}

func (r *NSCloneCommand) SetReceiversSpecifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCloneCommand) KeySpecifier() (*NSScriptObjectSpecifier, error) {
  return nil, fmt.Errorf("unimplemented")
}

