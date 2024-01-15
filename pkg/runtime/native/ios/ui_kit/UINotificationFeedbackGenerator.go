//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UINotificationFeedbackGenerator : UIKit.UIFeedbackGenerator
*/

type UINotificationFeedbackGenerator struct {
  *UIFeedbackGenerator

}

func (r *UINotificationFeedbackGenerator) NotificationOccurred() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

