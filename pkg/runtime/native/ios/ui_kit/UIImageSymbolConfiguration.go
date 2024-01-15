//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIImageSymbolConfiguration : UIKit.UIImageConfiguration
*/

type UIImageSymbolConfiguration struct {
  *UIImageConfiguration

}

func (r *UIImageSymbolConfiguration) ConfigurationWithPaletteColors() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageSymbolConfiguration) ConfigurationWithPointSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageSymbolConfiguration) ConfigurationWithFont() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageSymbolConfiguration) ConfigurationPreferringMulticolor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageSymbolConfiguration) ConfigurationPreferringMonochrome() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageSymbolConfiguration) ConfigurationWithoutTextStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageSymbolConfiguration) ConfigurationWithoutPointSizeAndWeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageSymbolConfiguration) ConfigurationWithScale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageSymbolConfiguration) ConfigurationWithHierarchicalColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageSymbolConfiguration) ConfigurationWithoutWeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageSymbolConfiguration) IsEqualToConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageSymbolConfiguration) ConfigurationWithWeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageSymbolConfiguration) ConfigurationWithTextStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageSymbolConfiguration) ConfigurationWithoutScale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageSymbolConfiguration) UnspecifiedConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

