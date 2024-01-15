//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIGraphicsImageRendererFormat : UIKit.UIGraphicsRendererFormat
*/

type UIGraphicsImageRendererFormat struct {
  *UIGraphicsRendererFormat

}

func (r *UIGraphicsImageRendererFormat) SetPreferredRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsImageRendererFormat) Scale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsImageRendererFormat) SupportsHighDynamicRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsImageRendererFormat) Opaque() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsImageRendererFormat) SetOpaque() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsImageRendererFormat) PrefersExtendedRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsImageRendererFormat) SetPrefersExtendedRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsImageRendererFormat) PreferredRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsImageRendererFormat) FormatForTraitCollection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsImageRendererFormat) SetScale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

