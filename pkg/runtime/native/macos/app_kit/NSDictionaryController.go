//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSDictionaryController : AppKit.NSArrayController
*/

type NSDictionaryController struct {
  *NSArrayController

}

func (r *NSDictionaryController) SetIncludedKeys() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDictionaryController) ExcludedKeys() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDictionaryController) LocalizedKeyDictionary() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDictionaryController) SetLocalizedKeyTable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDictionaryController) NewObject() (*NSDictionaryControllerKeyValuePair, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDictionaryController) IncludedKeys() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDictionaryController) SetExcludedKeys() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDictionaryController) SetLocalizedKeyDictionary() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDictionaryController) LocalizedKeyTable() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDictionaryController) SetInitialKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDictionaryController) InitialValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDictionaryController) SetInitialValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDictionaryController) InitialKey() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

