//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSScriptCommandDescription : objc.NSObject
*/

type NSScriptCommandDescription struct {
  *objc.NSObject
  *NSCoding
}

func (r *NSScriptCommandDescription) InitWithSuiteName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommandDescription) TypeForArgumentWithName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommandDescription) CreateCommandInstanceWithZone() (*NSScriptCommand, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommandDescription) CommandClassName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommandDescription) ArgumentNames() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommandDescription) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommandDescription) SuiteName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommandDescription) IsOptionalArgumentWithName() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommandDescription) CommandName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommandDescription) AppleEventClassCode() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommandDescription) AppleEventCodeForReturnType() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommandDescription) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommandDescription) AppleEventCodeForArgumentWithName() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommandDescription) CreateCommandInstance() (*NSScriptCommand, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommandDescription) AppleEventCode() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScriptCommandDescription) ReturnType() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

