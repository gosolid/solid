//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSCloseCommand : Foundation.NSScriptCommand
*/

type NSCloseCommand struct {
  *NSScriptCommand

}

func (r *NSCloseCommand) SaveOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

