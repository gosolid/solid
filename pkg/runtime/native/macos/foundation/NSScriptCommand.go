//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSScriptCommand : objc.NSObject
*/

type NSScriptCommand struct {
  *objc.NSObject
  *NSCoding
}

func (r *NSScriptCommand) Arguments() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) ScriptErrorNumber() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) SetScriptErrorOffendingObjectDescriptor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) SetScriptErrorExpectedTypeDescriptor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) CurrentCommand() (*NSScriptCommand, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) CommandDescription() (*NSScriptCommandDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) DirectParameter() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) PerformDefaultImplementation() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) EvaluatedArguments() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) ResumeExecutionWithResult() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) EvaluatedReceivers() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) SetArguments() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) SetScriptErrorNumber() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) ScriptErrorOffendingObjectDescriptor() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) ScriptErrorExpectedTypeDescriptor() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) ScriptErrorString() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) ReceiversSpecifier() (*NSScriptObjectSpecifier, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) SetScriptErrorString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) AppleEvent() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) SetDirectParameter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) ExecuteCommand() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) SuspendExecution() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) SetReceiversSpecifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) IsWellFormed() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommand) InitWithCommandDescription() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

