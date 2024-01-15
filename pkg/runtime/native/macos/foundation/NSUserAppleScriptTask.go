//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUserAppleScriptTask : Foundation.NSUserScriptTask
*/

type NSUserAppleScriptTask struct {
  *NSUserScriptTask

}

func (r *NSUserAppleScriptTask) ExecuteWithAppleEvent() error {
  return fmt.Errorf("unimplemented")
}

