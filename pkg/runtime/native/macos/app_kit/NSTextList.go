//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSTextList : objc.NSObject
*/

type NSTextList struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *NSTextList) IsOrdered() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextList) InitWithMarkerFormat() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextList) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextList) MarkerForItemNumber() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextList) MarkerFormat() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextList) ListOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextList) StartingItemNumber() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextList) SetStartingItemNumber() error {
  return fmt.Errorf("unimplemented")
}

