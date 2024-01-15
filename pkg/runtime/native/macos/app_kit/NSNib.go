//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSNib : objc.NSObject
*/

type NSNib struct {
  *objc.NSObject
  *foundation.NSCoding
}

func (r *NSNib) InitWithNibData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNib) InstantiateWithOwner() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSNib) InitWithNibNamed() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

