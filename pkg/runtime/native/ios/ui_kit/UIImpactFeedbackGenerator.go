//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIImpactFeedbackGenerator : UIKit.UIFeedbackGenerator
*/

type UIImpactFeedbackGenerator struct {
  *UIFeedbackGenerator

}

func (r *UIImpactFeedbackGenerator) InitWithStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImpactFeedbackGenerator) ImpactOccurred() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImpactFeedbackGenerator) ImpactOccurredWithIntensity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

