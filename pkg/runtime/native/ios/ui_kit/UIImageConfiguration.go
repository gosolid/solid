//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIImageConfiguration : objc.NSObject
*/

type UIImageConfiguration struct {
  *objc.NSObject

}

func (r *UIImageConfiguration) Locale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageConfiguration) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageConfiguration) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageConfiguration) ConfigurationWithTraitCollection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageConfiguration) ConfigurationWithLocale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageConfiguration) ConfigurationByApplyingConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageConfiguration) TraitCollection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

