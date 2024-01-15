//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UISelectionFeedbackGenerator : UIKit.UIFeedbackGenerator
*/

type UISelectionFeedbackGenerator struct {
  *UIFeedbackGenerator

}

func (r *UISelectionFeedbackGenerator) SelectionChanged() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

