//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSPICTImageRep : AppKit.NSImageRep
*/

type NSPICTImageRep struct {
  *NSImageRep

}

func (r *NSPICTImageRep) ImageRepWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPICTImageRep) InitWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPICTImageRep) PICTRepresentation() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPICTImageRep) BoundingBox() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

