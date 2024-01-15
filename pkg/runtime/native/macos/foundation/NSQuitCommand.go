//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSQuitCommand : Foundation.NSScriptCommand
*/

type NSQuitCommand struct {
  *NSScriptCommand

}

func (r *NSQuitCommand) SaveOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

