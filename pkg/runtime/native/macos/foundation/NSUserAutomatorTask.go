//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUserAutomatorTask : Foundation.NSUserScriptTask
*/

type NSUserAutomatorTask struct {
  *NSUserScriptTask

}

func (r *NSUserAutomatorTask) ExecuteWithInput() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserAutomatorTask) Variables() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserAutomatorTask) SetVariables() error {
  return fmt.Errorf("unimplemented")
}

