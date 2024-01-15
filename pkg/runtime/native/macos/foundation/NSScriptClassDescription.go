//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSScriptClassDescription : Foundation.NSClassDescription
*/

type NSScriptClassDescription struct {
  *NSClassDescription

}

func (r *NSScriptClassDescription) ClassDescriptionForClass() (*NSScriptClassDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptClassDescription) TypeForKey() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptClassDescription) ClassDescriptionForKey() (*NSScriptClassDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptClassDescription) AppleEventCodeForKey() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScriptClassDescription) HasOrderedToManyRelationshipForKey() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScriptClassDescription) HasReadablePropertyForKey() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScriptClassDescription) SuiteName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptClassDescription) AppleEventCode() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScriptClassDescription) DefaultSubcontainerAttributeKey() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptClassDescription) InitWithSuiteName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptClassDescription) IsLocationRequiredToCreateForKey() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScriptClassDescription) ImplementationClassName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptClassDescription) SuperclassDescription() (*NSScriptClassDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptClassDescription) MatchesAppleEventCode() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScriptClassDescription) SupportsCommand() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScriptClassDescription) SelectorForCommand() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptClassDescription) KeyWithAppleEventCode() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptClassDescription) HasPropertyForKey() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScriptClassDescription) HasWritablePropertyForKey() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScriptClassDescription) ClassName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

