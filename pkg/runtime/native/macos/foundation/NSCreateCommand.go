//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSCreateCommand : Foundation.NSScriptCommand
*/

type NSCreateCommand struct {
  *NSScriptCommand

}

func (r *NSCreateCommand) CreateClassDescription() (*NSScriptClassDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCreateCommand) ResolvedKeyDictionary() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

