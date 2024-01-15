//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIFontMetrics : objc.NSObject
*/

type UIFontMetrics struct {
  *objc.NSObject

}

func (r *UIFontMetrics) InitForTextStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontMetrics) ScaledFontForFont() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontMetrics) ScaledValueForValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontMetrics) DefaultMetrics() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontMetrics) MetricsForTextStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontMetrics) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

