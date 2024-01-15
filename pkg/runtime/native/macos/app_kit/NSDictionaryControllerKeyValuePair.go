//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSDictionaryControllerKeyValuePair : objc.NSObject
*/

type NSDictionaryControllerKeyValuePair struct {
  *objc.NSObject

}

func (r *NSDictionaryControllerKeyValuePair) SetValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDictionaryControllerKeyValuePair) LocalizedKey() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDictionaryControllerKeyValuePair) SetLocalizedKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDictionaryControllerKeyValuePair) IsExplicitlyIncluded() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDictionaryControllerKeyValuePair) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDictionaryControllerKeyValuePair) Key() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDictionaryControllerKeyValuePair) SetKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDictionaryControllerKeyValuePair) Value() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

