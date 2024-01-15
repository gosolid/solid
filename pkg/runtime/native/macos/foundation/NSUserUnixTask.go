//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUserUnixTask : Foundation.NSUserScriptTask
*/

type NSUserUnixTask struct {
  *NSUserScriptTask

}

func (r *NSUserUnixTask) ExecuteWithArguments() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserUnixTask) StandardInput() (*NSFileHandle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserUnixTask) SetStandardInput() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserUnixTask) StandardOutput() (*NSFileHandle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserUnixTask) SetStandardOutput() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserUnixTask) StandardError() (*NSFileHandle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserUnixTask) SetStandardError() error {
  return fmt.Errorf("unimplemented")
}

