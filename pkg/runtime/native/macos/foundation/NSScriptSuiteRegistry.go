//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSScriptSuiteRegistry : objc.NSObject
*/

type NSScriptSuiteRegistry struct {
  *objc.NSObject

}

func (r *NSScriptSuiteRegistry) SharedScriptSuiteRegistry() (*NSScriptSuiteRegistry, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptSuiteRegistry) LoadSuitesFromBundle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptSuiteRegistry) RegisterClassDescription() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptSuiteRegistry) AppleEventCodeForSuite() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScriptSuiteRegistry) SuiteNames() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptSuiteRegistry) SetSharedScriptSuiteRegistry() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptSuiteRegistry) LoadSuiteWithDictionary() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptSuiteRegistry) RegisterCommandDescription() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScriptSuiteRegistry) BundleForSuite() (*NSBundle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptSuiteRegistry) ClassDescriptionsInSuite() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptSuiteRegistry) CommandDescriptionsInSuite() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptSuiteRegistry) SuiteForAppleEventCode() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptSuiteRegistry) ClassDescriptionWithAppleEventCode() (*NSScriptClassDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptSuiteRegistry) CommandDescriptionWithAppleEventClass() (*NSScriptCommandDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptSuiteRegistry) AeteResource() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

