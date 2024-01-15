//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSTintConfiguration : objc.NSObject
*/

type NSTintConfiguration struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *NSTintConfiguration) BaseTintColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTintConfiguration) EquivalentContentTintColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTintConfiguration) AdaptsToUserAccentColor() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTintConfiguration) TintConfigurationWithPreferredColor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTintConfiguration) TintConfigurationWithFixedColor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTintConfiguration) DefaultTintConfiguration() (*NSTintConfiguration, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTintConfiguration) MonochromeTintConfiguration() (*NSTintConfiguration, error) {
  return nil, fmt.Errorf("unimplemented")
}

