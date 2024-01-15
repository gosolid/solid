//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSAppearance : objc.NSObject
*/

type NSAppearance struct {
  *objc.NSObject
  *foundation.NSSecureCoding
}

func (r *NSAppearance) PerformAsCurrentDrawingAppearance() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAppearance) AppearanceNamed() (*NSAppearance, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppearance) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppearance) BestMatchFromAppearancesWithNames() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppearance) SetCurrentAppearance() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAppearance) CurrentDrawingAppearance() (*NSAppearance, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppearance) AllowsVibrancy() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAppearance) InitWithAppearanceNamed() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppearance) Name() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppearance) CurrentAppearance() (*NSAppearance, error) {
  return nil, fmt.Errorf("unimplemented")
}

