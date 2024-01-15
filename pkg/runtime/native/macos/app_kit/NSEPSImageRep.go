//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSEPSImageRep : AppKit.NSImageRep
*/

type NSEPSImageRep struct {
  *NSImageRep

}

func (r *NSEPSImageRep) PrepareGState() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEPSImageRep) BoundingBox() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSEPSImageRep) EPSRepresentation() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEPSImageRep) ImageRepWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEPSImageRep) InitWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

