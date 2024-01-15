//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSScriptObjectSpecifier : objc.NSObject
*/

type NSScriptObjectSpecifier struct {
  *objc.NSObject
  *NSCoding
}

func (r *NSScriptObjectSpecifier) SetContainerIsObjectBeingTested() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) Key() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) SetKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) KeyClassDescription() (*NSScriptClassDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) SetEvaluationErrorNumber() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) ContainerIsRangeContainerObject() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) SetContainerIsRangeContainerObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) ContainerClassDescription() (*NSScriptClassDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) ObjectSpecifierWithDescriptor() (*NSScriptObjectSpecifier, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) InitWithContainerClassDescription() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) ChildSpecifier() (*NSScriptObjectSpecifier, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) SetContainerSpecifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) ObjectsByEvaluatingSpecifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) EvaluationErrorNumber() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) Descriptor() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) InitWithContainerSpecifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) SetChildSpecifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) SetContainerClassDescription() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) IndicesOfObjectsByEvaluatingWithContainer() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) ObjectsByEvaluatingWithContainers() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) ContainerSpecifier() (*NSScriptObjectSpecifier, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) ContainerIsObjectBeingTested() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScriptObjectSpecifier) EvaluationErrorSpecifier() (*NSScriptObjectSpecifier, error) {
  return nil, fmt.Errorf("unimplemented")
}

