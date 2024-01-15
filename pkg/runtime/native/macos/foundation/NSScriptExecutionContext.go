//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSScriptExecutionContext : objc.NSObject
*/

type NSScriptExecutionContext struct {
  *objc.NSObject

}

func (r *NSScriptExecutionContext) TopLevelObject() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptExecutionContext) SetTopLevelObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptExecutionContext) ObjectBeingTested() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptExecutionContext) SetObjectBeingTested() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptExecutionContext) RangeContainerObject() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptExecutionContext) SetRangeContainerObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptExecutionContext) SharedScriptExecutionContext() (*NSScriptExecutionContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

