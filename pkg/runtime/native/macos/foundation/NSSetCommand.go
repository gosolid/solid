//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSSetCommand : Foundation.NSScriptCommand
*/

type NSSetCommand struct {
  *NSScriptCommand

}

func (r *NSSetCommand) SetReceiversSpecifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSetCommand) KeySpecifier() (*NSScriptObjectSpecifier, error) {
  return nil, fmt.Errorf("unimplemented")
}

