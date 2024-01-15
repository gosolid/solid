//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIContentUnavailableTextProperties : objc.NSObject
*/

type UIContentUnavailableTextProperties struct {
  *objc.NSObject

}

func (r *UIContentUnavailableTextProperties) LineBreakMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableTextProperties) SetLineBreakMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableTextProperties) SetNumberOfLines() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableTextProperties) MinimumScaleFactor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableTextProperties) Font() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableTextProperties) SetFont() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableTextProperties) SetColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableTextProperties) SetMinimumScaleFactor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableTextProperties) AllowsDefaultTighteningForTruncation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableTextProperties) SetAllowsDefaultTighteningForTruncation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableTextProperties) Color() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableTextProperties) NumberOfLines() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableTextProperties) AdjustsFontSizeToFitWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableTextProperties) SetAdjustsFontSizeToFitWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

