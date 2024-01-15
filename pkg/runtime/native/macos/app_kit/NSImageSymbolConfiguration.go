//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSImageSymbolConfiguration : objc.NSObject
*/

type NSImageSymbolConfiguration struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *NSImageSymbolConfiguration) ConfigurationWithScale() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageSymbolConfiguration) ConfigurationPreferringMonochrome() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageSymbolConfiguration) ConfigurationPreferringHierarchical() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageSymbolConfiguration) ConfigurationWithPaletteColors() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageSymbolConfiguration) ConfigurationByApplyingConfiguration() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageSymbolConfiguration) ConfigurationWithTextStyle() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageSymbolConfiguration) ConfigurationWithHierarchicalColor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageSymbolConfiguration) ConfigurationPreferringMulticolor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageSymbolConfiguration) ConfigurationWithPointSize() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

