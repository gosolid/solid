//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSColorList : objc.NSObject
*/

type NSColorList struct {
  *objc.NSObject
  *foundation.NSSecureCoding
}

func (r *NSColorList) ColorListNamed() (*NSColorList, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorList) InitWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorList) ColorWithKey() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorList) SetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorList) RemoveColorWithKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorList) AllKeys() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorList) IsEditable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSColorList) InsertColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorList) WriteToURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSColorList) AvailableColorLists() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorList) WriteToFile() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSColorList) RemoveFile() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorList) Name() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

