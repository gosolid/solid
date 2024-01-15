//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSAppearanceCustomization
*/

type NSAppearanceCustomization struct {

}

func (r *NSAppearanceCustomization) Appearance() (*NSAppearance, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppearanceCustomization) SetAppearance() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAppearanceCustomization) EffectiveAppearance() (*NSAppearance, error) {
  return nil, fmt.Errorf("unimplemented")
}

