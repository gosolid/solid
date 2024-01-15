//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSTask : objc.NSObject
*/

type NSTask struct {
  *objc.NSObject

}

func (r *NSTask) TerminationHandler() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTask) Arguments() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTask) StandardInput() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTask) SetStandardInput() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTask) StandardOutput() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTask) StandardError() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTask) SetStandardError() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTask) TerminationStatus() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTask) SetExecutableURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTask) TerminationReason() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTask) Resume() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTask) CurrentDirectoryURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTask) SetQualityOfService() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTask) Interrupt() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTask) Suspend() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTask) SetCurrentDirectoryURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTask) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTask) SetArguments() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTask) SetEnvironment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTask) ProcessIdentifier() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTask) Terminate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTask) SetTerminationHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTask) Environment() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTask) ExecutableURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTask) SetStandardOutput() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTask) LaunchAndReturnError() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTask) QualityOfService() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTask) IsRunning() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

